// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package templ

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Header() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header><div class=\"navbar\"><div><a hx-get=\"/home\" hx-target=\"#content\" hx-push-url=\"true\"><img src=\"css/logo-nav.png\" class=\"logo\" height=\"60\" width=\"auto\" alt=\"\"></a></div><nav x-data=\"{ open: false }\" class=\"nav\"><button @click=\"open = !open\" :class=\"{ &#39;open&#39;: open }\" class=\"burger-menu\"><div class=\"burger-line line1\"></div><div class=\"burger-line line2\"></div><div class=\"burger-line line3\"></div></button><ul x-show=\"open\" x-transition:enter.duration.1100ms x-transition:enter.scale.50 x-transition:leave..duration.400ms x-transition:leave.scale.90 x-transition.scale.origin.top.right :class=\"{ &#39;open&#39; : open }\" @click.outside=\"open = false\" class=\"nav-items\" role=\"list\"><li class=\"nav-link\"><a class=\"button\" data-button-variant=\"anchor\" hx-get=\"/about\" hx-target=\"#content\" hx-swap=\"innerHTML\" hx-push-url=\"true\">About Me</a></li><li class=\"nav-link\"><a class=\"button\" data-button-variant=\"anchor\" hx-get=\"/blog\" hx-target=\"#content\" hx-swap=\"innerHTML \" hx-push-url=\"true\">Blog</a></li><li class=\"nav-link\"><a class=\"button\" data-button-variant=\"anchor\" hx-get=\"/contact\" hx-target=\"#content\" hx-swap=\"innerHTML \" hx-push-url=\"true\">Contact</a></li></ul><div :class=\"{ &#39;open&#39;: open }\" class=\"overlay\" @click=\"open = false\"></div></nav></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
