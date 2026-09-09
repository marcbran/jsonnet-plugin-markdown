package markdown

import (
	"github.com/google/go-jsonnet"
	"github.com/marcbran/jpoet/pkg/jpoet"
)

func Plugin(opts ...jpoet.PluginOption) *jpoet.Plugin {
	return jpoet.NewPlugin("markdown", []jsonnet.NativeFunction{
		ParseMarkdown(),
		ManifestMarkdown(),
	}, opts...)
}
