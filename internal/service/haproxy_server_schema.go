package service

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-opnsense/internal/tools"
)

// HaproxyServerResourceModel describes the resource data model.
type HaproxyServerResourceModel struct {
	Address              types.String `tfsdk:"address"`
	Advanced             types.String `tfsdk:"advanced"`
	CheckDownInterval    types.Int64  `tfsdk:"check_down_interval"`
	CheckInterval        types.Int64  `tfsdk:"check_interval"`
	Checkport            types.Int64  `tfsdk:"checkport"`
	Description          types.String `tfsdk:"description"`
	Enabled              types.Bool   `tfsdk:"enabled"`
	LinkedResolver       types.String `tfsdk:"linked_resolver"`
	MaxConnections       types.String `tfsdk:"max_connections"`
	Mode                 types.String `tfsdk:"mode"`
	MultiplexerProtocol  types.String `tfsdk:"multiplexer_protocol"`
	Name                 types.String `tfsdk:"name"`
	Number               types.String `tfsdk:"number"`
	Port                 types.Int64  `tfsdk:"port"`
	ResolvePrefer        types.String `tfsdk:"resolve_prefer"`
	ResolverOpts         types.String `tfsdk:"resolver_opts"`
	ServiceName          types.String `tfsdk:"service_name"`
	Source               types.String `tfsdk:"source"`
	Ssl                  types.Bool   `tfsdk:"ssl"`
	SslCA                types.String `tfsdk:"ssl_ca"`
	SslClientCertificate types.String `tfsdk:"ssl_client_certificate"`
	SslCRL               types.String `tfsdk:"ssl_crl"`
	SslSNI               types.String `tfsdk:"ssl_sni"`
	SslVerify            types.Bool   `tfsdk:"ssl_verify"`
	Type                 types.String `tfsdk:"type"`
	UnixSocket           types.String `tfsdk:"unix_socket"`
	Weight               types.String `tfsdk:"weight"`

	Id types.String `tfsdk:"id"`
}

func haproxyServerResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Host servers can be used to change DNS results from client queries or to add custom DNS records.",

		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"advanced": schema.StringAttribute{
				MarkdownDescription: "advanced",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"check_down_interval": schema.Int64Attribute{
				MarkdownDescription: "check_down_interval",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
			"check_interval": schema.Int64Attribute{
				MarkdownDescription: "basicAuthGroups",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
			},
			"checkport": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"max_connections": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("active"),
			},
			"multiplexer_protocol": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
			},
			"resolve_prefer": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"resolver_opts": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"service_name": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_ca": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_client_certificate": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_crl": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_sni": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_verify": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("static"),
			},
			"unix_socket": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"weight": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the host server.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func HaproxyServerDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Host servers can be used to change DNS results from client queries or to add custom DNS records.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"address": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"advanced": schema.StringAttribute{
				MarkdownDescription: "advanced",
				Computed:            true,
			},
			"check_down_interval": schema.Int64Attribute{
				MarkdownDescription: "check_down_interval",
				Computed:            true,
			},
			"check_interval": schema.Int64Attribute{
				MarkdownDescription: "basicAuthGroups",
				Computed:            true,
			},
			"checkport": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"max_connections": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"multiplexer_protocol": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"resolve_prefer": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"resolver_opts": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"service_name": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_ca": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_client_certificate": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_crl": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_sni": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_verify": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"unix_socket": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"weight": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
		},
	}
}

func convertHaproxyServerSchemaToStruct(d *HaproxyServerResourceModel) (*haproxy.Server, error) {
	return &haproxy.Server{
		Address:              d.Address.ValueString(),
		Advanced:             d.Advanced.ValueString(),
		CheckDownInterval:    tools.Int64ToString(d.CheckDownInterval.ValueInt64()),
		CheckInterval:        tools.Int64ToString(d.CheckInterval.ValueInt64()),
		Checkport:            tools.Int64ToString(d.Checkport.ValueInt64()),
		Description:          d.Description.ValueString(),
		Enabled:              tools.BoolToString(d.Enabled.ValueBool()),
		LinkedResolver:       api.SelectedMap(d.LinkedResolver.ValueString()),
		MaxConnections:       d.MaxConnections.ValueString(),
		Mode:                 api.SelectedMap(d.Mode.ValueString()),
		MultiplexerProtocol:  api.SelectedMap(d.MultiplexerProtocol.ValueString()),
		Name:                 d.Name.ValueString(),
		Number:               d.Number.ValueString(),
		Port:                 tools.Int64ToString(d.Port.ValueInt64()),
		ResolvePrefer:        api.SelectedMap(d.ResolvePrefer.ValueString()),
		ResolverOpts:         d.ResolverOpts.ValueString(),
		ServiceName:          d.ServiceName.ValueString(),
		Source:               d.Source.ValueString(),
		Ssl:                  tools.BoolToString(d.Ssl.ValueBool()),
		SslCA:                api.SelectedMap(d.SslCA.ValueString()),
		SslClientCertificate: api.SelectedMap(d.SslClientCertificate.ValueString()),
		SslCRL:               api.SelectedMap(d.SslCRL.ValueString()),
		SslSNI:               d.SslSNI.ValueString(),
		SslVerify:            tools.BoolToString(d.SslVerify.ValueBool()),
		Type:                 api.SelectedMap(d.Type.ValueString()),
		UnixSocket:           d.UnixSocket.ValueString(),
		Weight:               d.Weight.ValueString(),
	}, nil
}

func convertHaproxyServerStructToSchema(d *haproxy.Server) (*HaproxyServerResourceModel, error) {
	return &HaproxyServerResourceModel{
		Address:              types.StringValue(d.Address),
		Advanced:             types.StringValue(d.Advanced),
		CheckDownInterval:    types.Int64Value(tools.StringToInt64(d.CheckDownInterval)),
		CheckInterval:        types.Int64Value(tools.StringToInt64(d.CheckInterval)),
		Checkport:            types.Int64Value(tools.StringToInt64(d.Checkport)),
		Description:          types.StringValue(d.Description),
		Enabled:              types.BoolValue(tools.StringToBool(d.Enabled)),
		LinkedResolver:       types.StringValue(d.LinkedResolver.String()),
		MaxConnections:       types.StringValue(d.MaxConnections),
		Mode:                 types.StringValue(d.Mode.String()),
		MultiplexerProtocol:  types.StringValue(d.MultiplexerProtocol.String()),
		Name:                 types.StringValue(d.Name),
		Number:               types.StringValue(d.Number),
		Port:                 types.Int64Value(tools.StringToInt64(d.Port)),
		ResolvePrefer:        types.StringValue(d.ResolvePrefer.String()),
		ResolverOpts:         types.StringValue(d.ResolverOpts),
		ServiceName:          types.StringValue(d.ServiceName),
		Source:               types.StringValue(d.Source),
		Ssl:                  types.BoolValue(tools.StringToBool(d.Ssl)),
		SslCA:                types.StringValue(d.SslCA.String()),
		SslClientCertificate: types.StringValue(d.SslClientCertificate.String()),
		SslCRL:               types.StringValue(d.SslCRL.String()),
		SslSNI:               types.StringValue(d.SslSNI),
		SslVerify:            types.BoolValue(tools.StringToBool(d.SslVerify)),
		Type:                 types.StringValue(d.Type.String()),
		UnixSocket:           types.StringValue(d.UnixSocket),
		Weight:               types.StringValue(d.Weight),
	}, nil
}
