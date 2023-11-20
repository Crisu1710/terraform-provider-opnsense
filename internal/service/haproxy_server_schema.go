package service

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	MaxConnections       types.Int64  `tfsdk:"max_connections"`
	Mode                 types.String `tfsdk:"mode"`
	MultiplexerProtocol  types.String `tfsdk:"multiplexer_protocol"`
	Name                 types.String `tfsdk:"name"`
	Number               types.String `tfsdk:"number"`
	Port                 types.Int64  `tfsdk:"port"`
	ResolvePrefer        types.String `tfsdk:"resolve_prefer"`
	ResolverOpts         types.Set    `tfsdk:"resolver_opts"`
	ServiceName          types.String `tfsdk:"service_name"`
	Source               types.String `tfsdk:"source"`
	Ssl                  types.Bool   `tfsdk:"ssl"`
	SslCA                types.Set    `tfsdk:"ssl_ca"`
	SslClientCertificate types.String `tfsdk:"ssl_client_certificate"`
	SslCRL               types.String `tfsdk:"ssl_crl"`
	SslSNI               types.String `tfsdk:"ssl_sni"`
	SslVerify            types.Bool   `tfsdk:"ssl_verify"`
	Type                 types.String `tfsdk:"type"`
	UnixSocket           types.String `tfsdk:"unix_socket"`
	Weight               types.Int64  `tfsdk:"weight"`

	Id types.String `tfsdk:"id"`
}

func haproxyServerResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Host servers can be used to change DNS results from client queries or to add custom DNS records.",

		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				MarkdownDescription: "Provide either the FQDN or the IP address of the server.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"advanced": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"check_down_interval": schema.Int64Attribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on the server when the server state is DOWN. If it is not set HAProxy uses the check interval.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"check_interval": schema.Int64Attribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on this server. This setting takes precedence over default values in health monitors. It can still be overwritten from the backend pool.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"checkport": schema.Int64Attribute{
				MarkdownDescription: "Provide the TCP communication port to use during check.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for the server.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the server.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "This certificate will be sent if the server send a client certificate request.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"max_connections": schema.Int64Attribute{
				MarkdownDescription: "Specifies the maximal number of concurrent connections that will be sent to this server. The default value is 0 which means unlimited.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Sets the operation mode to use for this server.  (active, backup, disabled)",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "active", "backup", "disabled"),
				},
				Default: stringdefault.StaticString("active"),
			},
			"multiplexer_protocol": schema.StringAttribute{
				MarkdownDescription: "Forces the multiplexer's protocol to use for the outgoing connections to this server. (unspecified, fcgi, h1, h2)",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "unspecified", "fcgi", "h2", "h1"),
				},
				Default: stringdefault.StaticString("unspecified"),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify a static server. When creating a server template, then this prefix is used for the server names to be built.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "Provide the TCP or UDP communication port for this server.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"resolve_prefer": schema.StringAttribute{
				MarkdownDescription: "When DNS resolution is enabled for a server and multiple IP addresses from different families (ipv4, ipv6) are returned, HAProxy will prefer using an IP address from the selected family.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "ipv4", "ipv6"),
				},
				Default: stringdefault.StaticString(""),
			},
			"resolver_opts": schema.SetAttribute{
				MarkdownDescription: "Set latest allow-dup-ip, ignore-weight, prevent-dup-ip",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(stringvalidator.OneOf("", "allow-dup-ip", "ignore-weight", "prevent-dup-ip")),
				},
				Default: setdefault.StaticValue(tools.EmptySetValue()),
			},
			"service_name": schema.StringAttribute{
				MarkdownDescription: "Provide either the FQDN for all the servers this template initializes or a service name to discover the available services via DNS SRV records.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Sets the source address which will be used when connecting to the server.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable SSL communication with this server.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_ca": schema.SetAttribute{
				MarkdownDescription: "The selected CAs will be used to verify the server certificate.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"ssl_client_certificate": schema.StringAttribute{
				MarkdownDescription: "This certificate will be sent if the server send a client certificate request.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_crl": schema.StringAttribute{
				MarkdownDescription: "This certificate revocation list will be used to verify server's certificate.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_sni": schema.StringAttribute{
				MarkdownDescription: "The host name sent in the SNI TLS extension to the server.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_verify": schema.BoolAttribute{
				MarkdownDescription: "If disabled, server certificate is not verified. Otherwise the certificate provided by the server is verified using CAs and optional CRLs.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Either configure a static server or a template to initialize multiple servers with shared parameters. (static, template, unix)",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static", "template", "unix"),
				},
				Default: stringdefault.StaticString("static"),
			},
			"unix_socket": schema.StringAttribute{
				MarkdownDescription: "Select the frontend that provides the UNIX socket. This UNIX socket will be used as the server's address, making it possible to send connections to this frontend. Only frontends that provide the unix@ pattern as listen address can be selected.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"weight": schema.Int64Attribute{
				MarkdownDescription: "Adjust the server's weight relative to other servers.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
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
				MarkdownDescription: "Provide either the FQDN or the IP address of the server.",
				Computed:            true,
			},
			"advanced": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"check_down_interval": schema.Int64Attribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on the server when the server state is DOWN. If it is not set HAProxy uses the check interval.",
				Computed:            true,
			},
			"check_interval": schema.Int64Attribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on this server. This setting takes precedence over default values in health monitors. It can still be overwritten from the backend pool.",
				Computed:            true,
			},
			"checkport": schema.Int64Attribute{
				MarkdownDescription: "Provide the TCP communication port to use during check.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for the server.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable the server.",
				Computed:            true,
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "This certificate will be sent if the server send a client certificate request.",
				Computed:            true,
			},
			"max_connections": schema.Int64Attribute{
				MarkdownDescription: "Specifies the maximal number of concurrent connections that will be sent to this server. The default value is 0 which means unlimited.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Sets the operation mode to use for this server.  (active, backup, disabled)",
				Computed:            true,
			},
			"multiplexer_protocol": schema.StringAttribute{
				MarkdownDescription: "Forces the multiplexer's protocol to use for the outgoing connections to this server. (unspecified, fcgi, h1, h2)",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify a static server. When creating a server template, then this prefix is used for the server names to be built.",
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: "Provide the TCP or UDP communication port for this server.",
				Computed:            true,
			},
			"resolve_prefer": schema.StringAttribute{
				MarkdownDescription: "When DNS resolution is enabled for a server and multiple IP addresses from different families (ipv4, ipv6) are returned, HAProxy will prefer using an IP address from the selected family.",
				Computed:            true,
			},
			"resolver_opts": schema.SetAttribute{
				MarkdownDescription: "(allow-dup-ip, ignore-weight, prevent-dup-ip)",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"service_name": schema.StringAttribute{
				MarkdownDescription: "Provide either the FQDN for all the servers this template initializes or a service name to discover the available services via DNS SRV records.",
				Computed:            true,
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Sets the source address which will be used when connecting to the server.",
				Computed:            true,
			},
			"ssl": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable SSL communication with this server.",
				Computed:            true,
			},
			"ssl_ca": schema.SetAttribute{
				MarkdownDescription: "The selected CAs will be used to verify the server certificate.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"ssl_client_certificate": schema.StringAttribute{
				MarkdownDescription: "This certificate will be sent if the server send a client certificate request.",
				Computed:            true,
			},
			"ssl_crl": schema.StringAttribute{
				MarkdownDescription: "This certificate revocation list will be used to verify server's certificate.",
				Computed:            true,
			},
			"ssl_sni": schema.StringAttribute{
				MarkdownDescription: "The host name sent in the SNI TLS extension to the server.",
				Computed:            true,
			},
			"ssl_verify": schema.BoolAttribute{
				MarkdownDescription: "If disabled, server certificate is not verified. Otherwise the certificate provided by the server is verified using CAs and optional CRLs.",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Either configure a static server or a template to initialize multiple servers with shared parameters. (static, template, unix)",
				Computed:            true,
			},
			"unix_socket": schema.StringAttribute{
				MarkdownDescription: "Select the frontend that provides the UNIX socket. This UNIX socket will be used as the server's address, making it possible to send connections to this frontend. Only frontends that provide the unix@ pattern as listen address can be selected.",
				Computed:            true,
			},
			"weight": schema.Int64Attribute{
				MarkdownDescription: "Adjust the server's weight relative to other servers.",
				Computed:            true,
			},
		},
	}
}

func convertHaproxyServerSchemaToStruct(d *HaproxyServerResourceModel) (*haproxy.Server, error) {
	return &haproxy.Server{
		Address:              d.Address.ValueString(),
		Advanced:             d.Advanced.ValueString(),
		CheckDownInterval:    tools.Int64ToStringNegative(d.CheckDownInterval.ValueInt64()),
		CheckInterval:        tools.Int64ToStringNegative(d.CheckInterval.ValueInt64()),
		Checkport:            tools.Int64ToStringNegative(d.Checkport.ValueInt64()),
		Description:          d.Description.ValueString(),
		Enabled:              tools.BoolToString(d.Enabled.ValueBool()),
		LinkedResolver:       api.SelectedMap(d.LinkedResolver.ValueString()),
		MaxConnections:       tools.Int64ToStringNegative(d.MaxConnections.ValueInt64()),
		Mode:                 api.SelectedMap(d.Mode.ValueString()),
		MultiplexerProtocol:  api.SelectedMap(d.MultiplexerProtocol.ValueString()),
		Name:                 d.Name.ValueString(),
		Number:               d.Number.ValueString(),
		Port:                 tools.Int64ToStringNegative(d.Port.ValueInt64()),
		ResolvePrefer:        api.SelectedMap(d.ResolvePrefer.ValueString()),
		ResolverOpts:         tools.SetToString(d.ResolverOpts),
		ServiceName:          d.ServiceName.ValueString(),
		Source:               d.Source.ValueString(),
		Ssl:                  tools.BoolToString(d.Ssl.ValueBool()),
		SslCA:                tools.SetToString(d.SslCA),
		SslClientCertificate: api.SelectedMap(d.SslClientCertificate.ValueString()),
		SslCRL:               api.SelectedMap(d.SslCRL.ValueString()),
		SslSNI:               d.SslSNI.ValueString(),
		SslVerify:            tools.BoolToString(d.SslVerify.ValueBool()),
		Type:                 api.SelectedMap(d.Type.ValueString()),
		UnixSocket:           api.SelectedMap(d.UnixSocket.ValueString()),
		Weight:               tools.Int64ToStringNegative(d.Weight.ValueInt64()),
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
		MaxConnections:       types.Int64Value(tools.StringToInt64(d.MaxConnections)),
		Mode:                 types.StringValue(d.Mode.String()),
		MultiplexerProtocol:  types.StringValue(d.MultiplexerProtocol.String()),
		Name:                 types.StringValue(d.Name),
		Number:               types.StringValue(d.Number),
		Port:                 types.Int64Value(tools.StringToInt64(d.Port)),
		ResolvePrefer:        types.StringValue(d.ResolvePrefer.String()),
		ResolverOpts:         tools.StringToSet(d.ResolverOpts),
		ServiceName:          types.StringValue(d.ServiceName),
		Source:               types.StringValue(d.Source),
		Ssl:                  types.BoolValue(tools.StringToBool(d.Ssl)),
		SslCA:                tools.StringToSet(d.SslCA),
		SslClientCertificate: types.StringValue(d.SslClientCertificate.String()),
		SslCRL:               types.StringValue(d.SslCRL.String()),
		SslSNI:               types.StringValue(d.SslSNI),
		SslVerify:            types.BoolValue(tools.StringToBool(d.SslVerify)),
		Type:                 types.StringValue(d.Type.String()),
		UnixSocket:           types.StringValue(d.UnixSocket.String()),
		Weight:               types.Int64Value(tools.StringToInt64(d.Weight)),
	}, nil
}
