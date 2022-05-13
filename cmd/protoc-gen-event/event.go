package main

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/voi-oss/protoc-gen-event/pkg/options"
)

const (
	contextPkg   = protogen.GoImportPath("context")
	messagePkg   = protogen.GoImportPath("github.com/ThreeDotsLabs/watermill/message")
	watermillPkg = protogen.GoImportPath("github.com/ThreeDotsLabs/watermill")
	protoJSONPkg = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
)

type GeneratorConfig struct {
	Suffix string
}

type messageOptions struct {
	topic string
	skip  bool
}

type fieldOptions struct {
	isMetadata bool
	name       string
}

// generateFile generates a _grpc.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File, config GeneratorConfig) *protogen.GeneratedFile {
	if len(file.Messages) == 0 {
		return nil
	}

	generated := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+"_event.pb.go", file.GoImportPath)
	generated.P("// Code generated by protoc-gen-event. DO NOT EDIT.")
	generated.P()
	generated.P("package ", file.GoPackageName)
	generated.P()

	generateFileContent(file, generated, config)

	return generated
}

// generateFileContent generates events definitions, excluding the package statement.
func generateFileContent(file *protogen.File, g *protogen.GeneratedFile, config GeneratorConfig) {
	if len(file.Messages) == 0 {
		return
	}

	for _, message := range file.Messages {
		generateEvent(g, message, config)
	}
}

// generateEvent generates the functions needed to integrate an event with Watermill.
func generateEvent(g *protogen.GeneratedFile, message *protogen.Message, config GeneratorConfig) {
	opts := getOptions(message, config)
	if opts.skip {
		return
	}

	g.P("// Publish will JSON marshal and publish this on a publisher")
	g.Annotate(message.GoIdent.String()+".Publish", message.Location)
	g.P("func (e *", message.GoIdent, ") Publish(ctx ", g.QualifiedGoIdent(contextPkg.Ident("Context")), ", publisher ", g.QualifiedGoIdent(messagePkg.Ident("Publisher")), ") error {")
	g.P("return e.PublishWithUUID(ctx, publisher, ", watermillPkg.Ident("NewUUID"), "()", ")")
	g.P("}")

	g.P()

	g.P("// PublishWithUUID will JSON marshal and publish this on a publisher with the given UUID")
	g.Annotate(message.GoIdent.String()+".PublishWithUUID", message.Location)
	g.P("func (e *", message.GoIdent, ") PublishWithUUID(ctx ", g.QualifiedGoIdent(contextPkg.Ident("Context")), ", publisher ", g.QualifiedGoIdent(messagePkg.Ident("Publisher")), ", uuid string) error {")
	g.P("payload, err := ", g.QualifiedGoIdent(protoJSONPkg.Ident("Marshal")), "(x)")
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")

	g.P()

	g.P("msg := message.NewMessage(uuid, payload)")

	for _, field := range message.Fields {
		fo := getFieldOptions(field)
		if !fo.isMetadata {
			continue
		}
		g.P("msg.Metadata.Set(\"", fo.name, "\", e.Get", field.GoName, "())")
	}

	g.P("msg.SetContext(ctx)")
	g.P()

	g.P("return publisher.Publish(\"", opts.topic, "\", msg)")
	g.P("}")

	g.P()

	g.P("// Simple consumer wrapper")
	g.P("type int", message.GoIdent, "Handler struct {")
	g.P("HandlerFunc func(pe *", message.GoIdent, ", m *", messagePkg.Ident("Message"), ") error")
	g.P("}")

	g.P("func ", message.GoIdent, "Handler(f func(pe *", message.GoIdent, ", m *", messagePkg.Ident("Message"), ") error) *int", message.GoIdent, "Handler {")
	g.P("return &int", message.GoIdent, "Handler {")
	g.P("HandlerFunc: f,")
	g.P("}")
	g.P("}")

	g.P()
	g.P("func (h *int", message.GoIdent, "Handler) Topic() string { ")
	g.P("return \"", opts.topic, "\"")
	g.P("}")

	g.P()
	g.P("func (h *int", message.GoIdent, "Handler) Handle(m *", messagePkg.Ident("Message"), ") error { ")
	g.P("pe := &", message.GoIdent, "{}")
	g.P("opts := protojson.UnmarshalOptions{")
	g.P("DiscardUnknown: true,")
	g.P("}")
	g.P("if err := opts.Unmarshal(m.Payload, pe); err != nil {")
	g.P("return err")
	g.P("}")
	g.P()
	g.P("return h.HandlerFunc(pe, m)")
	g.P("}")
}

func getOptions(message *protogen.Message, config GeneratorConfig) messageOptions {
	mo := messageOptions{
		topic: string(message.Desc.FullName()),
		skip:  !strings.HasSuffix(string(message.Desc.Name()), config.Suffix),
	}

	msgOpts := message.Desc.Options().(*descriptorpb.MessageOptions)
	opts := proto.GetExtension(msgOpts, options.E_Options).(*options.EventOption)

	if opts.GetTopicName() != "" {
		mo.topic = opts.GetTopicName()
	}
	if opts.GetSkip() {
		mo.skip = opts.GetSkip()
	}

	return mo
}

func getFieldOptions(field *protogen.Field) fieldOptions {
	fo := fieldOptions{
		isMetadata: false,
		name:       field.GoName,
	}

	msgOpts := field.Desc.Options().(*descriptorpb.FieldOptions)
	opts := proto.GetExtension(msgOpts, options.E_Field).(*options.FieldOption)

	if opts.GetIsMetadata() {
		fo.isMetadata = opts.GetIsMetadata()
	}
	if opts.GetName() != "" {
		fo.name = opts.GetName()
	}

	return fo
}
