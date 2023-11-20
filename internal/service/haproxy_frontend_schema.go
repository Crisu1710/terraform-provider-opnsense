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

// HaproxyFrontendResourceModel describes the resource data model.
type HaproxyFrontendResourceModel struct {
	AdvertisedProtocols          types.Set    `tfsdk:"advertised_protocols"`
	BasicAuthEnabled             types.Bool   `tfsdk:"basic_auth_enabled"`
	BasicAuthGroups              types.Set    `tfsdk:"basic_auth_groups"`
	BasicAuthUsers               types.Set    `tfsdk:"basic_auth_users"`
	Bind                         types.Set    `tfsdk:"bind"`
	BindOptions                  types.String `tfsdk:"bind_options"`
	ConnectionBehaviour          types.String `tfsdk:"connection_behaviour"`
	CustomOptions                types.String `tfsdk:"custom_options"`
	DefaultBackend               types.String `tfsdk:"default_backend"`
	Description                  types.String `tfsdk:"description"`
	Enabled                      types.Bool   `tfsdk:"enabled"`
	ForwardFor                   types.Bool   `tfsdk:"forward_for"`
	Http2Enabled                 types.Bool   `tfsdk:"http2enabled"`
	Http2EnabledNontls           types.Bool   `tfsdk:"http2enabled_nontls"`
	LinkedActions                types.Set    `tfsdk:"linked_actions"`
	LinkedCpuAffinityRules       types.Set    `tfsdk:"linked_cpu_affinity_rules"`
	LinkedErrorfiles             types.Set    `tfsdk:"linked_errorfiles"`
	LoggingDetailedLog           types.Bool   `tfsdk:"logging_detailed_log"`
	LoggingDontLogNormal         types.Bool   `tfsdk:"logging_dont_log_normal"`
	LoggingDontLogNull           types.Bool   `tfsdk:"logging_dont_log_null"`
	LoggingLogSeparateErrors     types.Bool   `tfsdk:"logging_log_separate_errors"`
	LoggingSocketStats           types.Bool   `tfsdk:"logging_socket_stats"`
	Mode                         types.String `tfsdk:"mode"`
	Name                         types.String `tfsdk:"name"`
	PrometheusEnabled            types.Bool   `tfsdk:"prometheus_enabled"`
	PrometheusPath               types.String `tfsdk:"prometheus_path"`
	SslAdvancedEnabled           types.Bool   `tfsdk:"ssl_advanced_enabled"`
	SslBindOptions               types.Set    `tfsdk:"ssl_bind_options"`
	SslCertificates              types.Set    `tfsdk:"ssl_certificates"`
	SslCipherList                types.String `tfsdk:"ssl_cipher_list"`
	SslCipherSuites              types.String `tfsdk:"ssl_cipher_suites"`
	SslClientAuthCAs             types.Set    `tfsdk:"ssl_client_auth_cas"`
	SslClientAuthCRLs            types.Set    `tfsdk:"ssl_client_auth_crls"`
	SslClientAuthEnabled         types.Bool   `tfsdk:"ssl_client_auth_enabled"`
	SslClientAuthVerify          types.String `tfsdk:"ssl_client_auth_verify"`
	SslCustomOptions             types.String `tfsdk:"ssl_custom_options"`
	SslDefaultCertificate        types.String `tfsdk:"ssl_default_certificate"`
	SslEnabled                   types.Bool   `tfsdk:"ssl_enabled"`
	SslHstsEnabled               types.Bool   `tfsdk:"ssl_hsts_enabled"`
	SslHstsIncludeSubDomains     types.Bool   `tfsdk:"ssl_hsts_include_sub_domains"`
	SslHstsMaxAge                types.Int64  `tfsdk:"ssl_hsts_max_age"`
	SslHstsPreload               types.Bool   `tfsdk:"ssl_hsts_preload"`
	SslMaxVersion                types.String `tfsdk:"ssl_max_version"`
	SslMinVersion                types.String `tfsdk:"ssl_min_version"`
	StickinessBytesInRatePeriod  types.String `tfsdk:"stickiness_bytes_in_rate_period"`
	StickinessBytesOutRatePeriod types.String `tfsdk:"stickiness_bytes_out_rate_period"`
	StickinessConnRatePeriod     types.String `tfsdk:"stickiness_conn_rate_period"`
	StickinessCounter            types.Bool   `tfsdk:"stickiness_counter"`
	StickinessCounterKey         types.String `tfsdk:"stickiness_counter_key"`
	StickinessDataTypes          types.Set    `tfsdk:"stickiness_data_types"`
	StickinessExpire             types.String `tfsdk:"stickiness_expire"`
	StickinessHttpErrRatePeriod  types.String `tfsdk:"stickiness_http_err_rate_period"`
	StickinessHttpReqRatePeriod  types.String `tfsdk:"stickiness_http_req_rate_period"`
	StickinessLength             types.Int64  `tfsdk:"stickiness_length"`
	StickinessPattern            types.String `tfsdk:"stickiness_pattern"`
	StickinessSessRatePeriod     types.String `tfsdk:"stickiness_sess_rate_period"`
	StickinessSize               types.String `tfsdk:"stickiness_size"`
	TuningMaxConnections         types.Int64  `tfsdk:"tuning_max_connections"`
	TuningShards                 types.String `tfsdk:"tuning_shards"`
	TuningTimeoutClient          types.String `tfsdk:"tuning_timeout_client"`
	TuningTimeoutHttpKeepAlive   types.String `tfsdk:"tuning_timeout_http_keep_alive"`
	TuningTimeoutHttpReq         types.String `tfsdk:"tuning_timeout_http_req"`

	Id types.String `tfsdk:"id"`
}

func haproxyFrontendResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Host frontends can be used to change DNS results from client queries or to add custom DNS records.",

		Attributes: map[string]schema.Attribute{
			"advertised_protocols": schema.SetAttribute{
				MarkdownDescription: "When using the TLS ALPN extension, HAProxy advertises the specified protocol list as supported on top of ALPN. SSL offloading must be enabled.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(stringvalidator.OneOf("h2", "http11", "http10"))},
				ElementType: types.StringType,
				Default:     setdefault.StaticValue(tools.AppendSetValue([]string{"h2", "http11"})),
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP Basic Authentication.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"basic_auth_groups": schema.SetAttribute{
				MarkdownDescription: "Allowed Groups",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"basic_auth_users": schema.SetAttribute{
				MarkdownDescription: "Allowed Users",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"bind": schema.SetAttribute{
				MarkdownDescription: "Configure listen addresses for this Frontend.",
				Required:            true,
				ElementType:         types.StringType,
			},
			"bind_options": schema.StringAttribute{
				MarkdownDescription: "A list of parameters that will be appended to every Listen Address line.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"connection_behaviour": schema.StringAttribute{
				MarkdownDescription: "By default HAProxy operates in keep-alive mode with regards to persistent connections. Option 'httpclose' configures HAProxy to close connections with the server and the client as soon as the request and the response are received. It will also check if a 'Connection: close' header is already set in each direction, and will add one if missing. Option 'http-server-close' enables HTTP connection-close mode on the server side while keeping the ability to support HTTP keep-alive and pipelining on the client side. ",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http-keep-alive", "httpclose", "http-server-close"),
				},
				Default: stringdefault.StaticString("http-keep-alive"),
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "These lines will be added to the HAProxy frontend configuration.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"default_backend": schema.StringAttribute{
				MarkdownDescription: "Set the default Backend Pool to use for this Frontend",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this Frontend",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this Frontend",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"forward_for": schema.BoolAttribute{
				MarkdownDescription: "Enable insertion of the X-Forwarded-For header to requests sent to servers.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http2enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable support for HTTP/2.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http2enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "Enable support for HTTP/2 even if TLS (SSL offloading) is not enabled.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"linked_actions": schema.SetAttribute{
				MarkdownDescription: "Choose rules to be included in the Frontend.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"linked_cpu_affinity_rules": schema.SetAttribute{
				MarkdownDescription: "Choose CPU affinity rules that should be applied to this Public Service.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"linked_errorfiles": schema.SetAttribute{
				MarkdownDescription: "Choose error messages to be included in the Frontend",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"logging_detailed_log": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable verbose logging. Each log line turns into a much richer format.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_dont_log_normal": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable logging of normal, successful connections.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_dont_log_null": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable logging of connections with no data.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_log_separate_errors": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable collecting & providing separate statistics for each socket.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_socket_stats": schema.BoolAttribute{
				MarkdownDescription: "Allow HAProxy to automatically raise log level for non-completely successful connections to aid debugging.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Set the running mode or protocol for this Frontend Available values: `http`, `ssl`, ``",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http", "ssl", "tcp"),
				},
				Default: stringdefault.StaticString("http"),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify the Frontend",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"prometheus_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HAProxy's Prometheus exporter.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"prometheus_path": schema.StringAttribute{
				MarkdownDescription: "The path where the Prometheus exporter can be accessed.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/metrics"),
			},
			"ssl_advanced_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable advanced SSL settings.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_bind_options": schema.SetAttribute{
				MarkdownDescription: "Used to enforce or disable certain SSL options.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(stringvalidator.OneOf("no-sslv3", "no-tlsv10", "no-tlsv11", "no-tlsv12", "no-tlsv13", "no-tls-tickets", "force-sslv3", "force-tlsv10", "force-tlsv11", "force-tlsv12", "force-tlsv13", "prefer-client-ciphers", "strict-sni")),
				},
				Default: setdefault.StaticValue(tools.EmptySetValue()),
			},
			"ssl_certificates": schema.SetAttribute{
				MarkdownDescription: "This certificate will be presented if no SNI is provided by the client or if the client provides an SNI hostname which does not match any certificate.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"ssl_cipher_list": schema.StringAttribute{
				MarkdownDescription: "It sets the default string describing the list of cipher algorithms (\"cipher suite\") that are negotiated during the SSL/TLS handshake up to TLSv1.2.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256"),
			},
			"ssl_cipher_suites": schema.StringAttribute{
				MarkdownDescription: "It sets the default string describing the list of cipher algorithms (cipher suite) that are negotiated during the SSL/TLS handshake for TLSv1.3.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("TLS_AES_128_GCM_SHA256:TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256"),
			},
			"ssl_client_auth_cas": schema.SetAttribute{
				MarkdownDescription: "Select CA certificates to use for client certificate authentication.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"ssl_client_auth_crls": schema.SetAttribute{
				MarkdownDescription: "Select CRLs to use for client certificate authentication.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"ssl_client_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable Client Certificate Authentication.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_client_auth_verify": schema.StringAttribute{
				MarkdownDescription: "If set to 'optional' or 'required', client certificate is requested.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "none", "optional", "required"),
				},
				Default: stringdefault.StaticString("required"),
			},
			"ssl_custom_options": schema.StringAttribute{
				MarkdownDescription: "Pass additional SSL parameters to the HAProxy configuration.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_default_certificate": schema.StringAttribute{
				MarkdownDescription: "This certificate will be presented if no SNI is provided by the client or if the client provides an SNI hostname which does not match any certificate. This parameter is optional to enforce a certain sort order for certificates.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable SSL offloading.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_hsts_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP Strict Transport Security.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ssl_hsts_include_sub_domains": schema.BoolAttribute{
				MarkdownDescription: "Enable if all present and future subdomains will be HTTPS.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_hsts_max_age": schema.Int64Attribute{
				MarkdownDescription: "Future requests to the domain should use only HTTPS for the specified time (in seconds).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(15768000),
			},
			"ssl_hsts_preload": schema.BoolAttribute{
				MarkdownDescription: "Enable if you like this domain to be included in the HSTS preload list.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_max_version": schema.StringAttribute{
				MarkdownDescription: "This option enforces use of the specified version (or lower) on SSL connections.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "SSLv3", "TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3"),
				},
				Default: stringdefault.StaticString(""),
			},
			"ssl_min_version": schema.StringAttribute{
				MarkdownDescription: "This option enforces use of the specified version (or higher) on SSL connections.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "SSLv3", "TLSv1.0", "TLSv1.1", "TLSv1.2", "TLSv1.3"),
				},
				Default: stringdefault.StaticString(""),
			},
			"stickiness_bytes_in_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"stickiness_bytes_out_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average outgoing bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"stickiness_conn_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming connection rate over that period, in connections per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_counter": schema.BoolAttribute{
				MarkdownDescription: "Enable to be able to retrieve values from sticky counters. If disabled, all values will return 0, rendering many conditions useless. Sticky counter key",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"stickiness_counter_key": schema.StringAttribute{
				MarkdownDescription: "It describes what elements of the incoming request or connection will be analyzed, extracted, combined, and used to select which table entry to update the counters. Defaults to 'src' to track elements of the source IP. See the [HAProxy documentation](http://docs.haproxy.org/2.6/configuration.html#tcp-request%20connection) for a full description.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("src"),
			},
			"stickiness_data_types": schema.SetAttribute{
				MarkdownDescription: "This is used to store additional information in the stick-table. It may be used by ACLs in order to control various criteria related to the activity of the client matching the stick-table. Note that this directly impacts memory usage. See the [HAProxy documentation](http://docs.haproxy.org/2.6/configuration.html#stick-table) for a full description.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(stringvalidator.OneOf("", "conn_cnt", "conn_cur", "conn_rate", "sess_cnt", "sess_rate", "http_req_cnt", "http_req_rate", "http_err_cnt", "http_err_rate", "bytes_in_cnt", "bytes_in_rate", "bytes_out_cnt", "bytes_out_rate")),
				},
				Default: setdefault.StaticValue(tools.EmptySetValue()),
			},
			"stickiness_expire": schema.StringAttribute{
				MarkdownDescription: "Enter a number followed by one of the supported suffixes d, h, m, s, ms. This configures the maximum duration of an entry in the stick-table since it was last created, refreshed or matched. The maximum duration is slightly above 24 days.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("30m"),
			},
			"stickiness_http_err_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request error rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_http_req_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_length": schema.Int64Attribute{
				MarkdownDescription: "Specify the maximum length for a value in the stick-table.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "Choose the type of data that should be stored in this stick-table.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "ipv4", "ipv6", "integer", "string", "binary"),
				},
				Default: stringdefault.StaticString(""),
			},
			"stickiness_sess_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming session rate over that period, in sessions per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_size": schema.StringAttribute{
				MarkdownDescription: "Enter a number followed by one of the supported suffixes k, m, g. This configures the maximum number of entries that can fit in the table. This value directly impacts memory usage. Count approximately 50 bytes per entry, plus the size of a string if any.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("50k"),
			},
			"tuning_max_connections": schema.Int64Attribute{
				MarkdownDescription: "Set the maximum number of concurrent connections for this Public Service.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"tuning_shards": schema.StringAttribute{
				MarkdownDescription: "This option automatically creates the specified number of listeners for every IP:port combination and evenly distributes them among available threads.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_client": schema.StringAttribute{
				MarkdownDescription: "Set the maximum inactivity time on the client side. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_http_keep_alive": schema.StringAttribute{
				MarkdownDescription: "Set the maximum allowed time to wait for a new HTTP request to appear. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_http_req": schema.StringAttribute{
				MarkdownDescription: "Set the maximum allowed time to wait for a complete HTTP request. In order to offer DoS protection, it may be required to lower the maximum accepted time to receive a complete HTTP request without affecting the client timeout. This helps protecting against established connections on which nothing is sent. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the host frontend.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func HaproxyFrontendDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Host frontends can be used to change DNS results from client queries or to add custom DNS records.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"advertised_protocols": schema.SetAttribute{
				MarkdownDescription: "When using the TLS ALPN extension, HAProxy advertises the specified protocol list as supported on top of ALPN. SSL offloading must be enabled.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP Basic Authentication.",
				Computed:            true,
			},
			"basic_auth_groups": schema.SetAttribute{
				MarkdownDescription: "Allowed Groups",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"basic_auth_users": schema.SetAttribute{
				MarkdownDescription: "Allowed Users",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"bind": schema.SetAttribute{
				MarkdownDescription: "Configure listen addresses for this Frontend.",
				Required:            true,
				ElementType:         types.StringType,
			},
			"bind_options": schema.StringAttribute{
				MarkdownDescription: "A list of parameters that will be appended to every Listen Address line.",
				Computed:            true,
			},
			"connection_behaviour": schema.StringAttribute{
				MarkdownDescription: "By default HAProxy operates in keep-alive mode with regards to persistent connections. Option 'httpclose' configures HAProxy to close connections with the server and the client as soon as the request and the response are received. It will also check if a 'Connection: close' header is already set in each direction, and will add one if missing. Option 'http-server-close' enables HTTP connection-close mode on the server side while keeping the ability to support HTTP keep-alive and pipelining on the client side. ",
				Computed:            true,
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "These lines will be added to the HAProxy frontend configuration.",
				Computed:            true,
			},
			"default_backend": schema.StringAttribute{
				MarkdownDescription: "Set the default Backend Pool to use for this Frontend",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this Frontend",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this Frontend",
				Computed:            true,
			},
			"forward_for": schema.BoolAttribute{
				MarkdownDescription: "Enable insertion of the X-Forwarded-For header to requests sent to servers.",
				Computed:            true,
			},
			"http2enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable support for HTTP/2.",
				Computed:            true,
			},
			"http2enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "Enable support for HTTP/2 even if TLS (SSL offloading) is not enabled.",
				Computed:            true,
			},
			"linked_actions": schema.SetAttribute{
				MarkdownDescription: "Choose rules to be included in the Frontend.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"linked_cpu_affinity_rules": schema.SetAttribute{
				MarkdownDescription: "Choose CPU affinity rules that should be applied to this Public Service.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"linked_errorfiles": schema.SetAttribute{
				MarkdownDescription: "Choose error messages to be included in the Frontend",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"logging_detailed_log": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable verbose logging. Each log line turns into a much richer format.",
				Computed:            true,
			},
			"logging_dont_log_normal": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable logging of normal, successful connections.",
				Computed:            true,
			},
			"logging_dont_log_null": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable logging of connections with no data.",
				Computed:            true,
			},
			"logging_log_separate_errors": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable collecting & providing separate statistics for each socket.",
				Computed:            true,
			},
			"logging_socket_stats": schema.BoolAttribute{
				MarkdownDescription: "Allow HAProxy to automatically raise log level for non-completely successful connections to aid debugging.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Set the running mode or protocol for this Frontend",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify the Frontend",
				Required:            true,
			},
			"prometheus_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HAProxy's Prometheus exporter.",
				Computed:            true,
			},
			"prometheus_path": schema.StringAttribute{
				MarkdownDescription: "The path where the Prometheus exporter can be accessed.",
				Computed:            true,
			},
			"ssl_advanced_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable advanced SSL settings.",
				Computed:            true,
			},
			"ssl_bind_options": schema.SetAttribute{
				MarkdownDescription: "Used to enforce or disable certain SSL options.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"ssl_certificates": schema.SetAttribute{
				MarkdownDescription: "This certificate will be presented if no SNI is provided by the client or if the client provides an SNI hostname which does not match any certificate.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"ssl_cipher_list": schema.StringAttribute{
				MarkdownDescription: "It sets the default string describing the list of cipher algorithms (\"cipher suite\") that are negotiated during the SSL/TLS handshake up to TLSv1.2.",
				Computed:            true,
			},
			"ssl_cipher_suites": schema.StringAttribute{
				MarkdownDescription: "It sets the default string describing the list of cipher algorithms (cipher suite) that are negotiated during the SSL/TLS handshake for TLSv1.3.",
				Computed:            true,
			},
			"ssl_client_auth_cas": schema.SetAttribute{
				MarkdownDescription: "Select CA certificates to use for client certificate authentication.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"ssl_client_auth_crls": schema.SetAttribute{
				MarkdownDescription: "Select CRLs to use for client certificate authentication.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"ssl_client_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable Client Certificate Authentication.",
				Computed:            true,
			},
			"ssl_client_auth_verify": schema.StringAttribute{
				MarkdownDescription: "If set to 'optional' or 'required', client certificate is requested.",
				Computed:            true,
			},
			"ssl_custom_options": schema.StringAttribute{
				MarkdownDescription: "Pass additional SSL parameters to the HAProxy configuration.",
				Computed:            true,
			},
			"ssl_default_certificate": schema.StringAttribute{
				MarkdownDescription: "This certificate will be presented if no SNI is provided by the client or if the client provides an SNI hostname which does not match any certificate. This parameter is optional to enforce a certain sort order for certificates.",
				Computed:            true,
			},
			"ssl_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable SSL offloading.",
				Computed:            true,
			},
			"ssl_hsts_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP Strict Transport Security.",
				Computed:            true,
			},
			"ssl_hsts_include_sub_domains": schema.BoolAttribute{
				MarkdownDescription: "Enable if all present and future subdomains will be HTTPS.",
				Computed:            true,
			},
			"ssl_hsts_max_age": schema.Int64Attribute{
				MarkdownDescription: "Future requests to the domain should use only HTTPS for the specified time (in seconds).",
				Computed:            true,
			},
			"ssl_hsts_preload": schema.BoolAttribute{
				MarkdownDescription: "Enable if you like this domain to be included in the HSTS preload list.",
				Computed:            true,
			},
			"ssl_max_version": schema.StringAttribute{
				MarkdownDescription: "This option enforces use of the specified version (or lower) on SSL connections.",
				Computed:            true,
			},
			"ssl_min_version": schema.StringAttribute{
				MarkdownDescription: "This option enforces use of the specified version (or higher) on SSL connections.",
				Computed:            true,
			},
			"stickiness_bytes_in_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"stickiness_bytes_out_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average outgoing bytes rate over that period, in bytes per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"stickiness_conn_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming connection rate over that period, in connections per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"stickiness_counter": schema.BoolAttribute{
				MarkdownDescription: "Enable to be able to retrieve values from sticky counters. If disabled, all values will return 0, rendering many conditions useless. Sticky counter key",
				Computed:            true,
			},
			"stickiness_counter_key": schema.StringAttribute{
				MarkdownDescription: "It describes what elements of the incoming request or connection will be analyzed, extracted, combined, and used to select which table entry to update the counters. Defaults to 'src' to track elements of the source IP. See the [HAProxy documentation](http://docs.haproxy.org/2.6/configuration.html#tcp-request%20connection) for a full description.",
				Computed:            true,
			},
			"stickiness_data_types": schema.SetAttribute{
				MarkdownDescription: "This is used to store additional information in the stick-table. It may be used by ACLs in order to control various criteria related to the activity of the client matching the stick-table. Note that this directly impacts memory usage. See the [HAProxy documentation](http://docs.haproxy.org/2.6/configuration.html#stick-table) for a full description.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"stickiness_expire": schema.StringAttribute{
				MarkdownDescription: "Enter a number followed by one of the supported suffixes d, h, m, s, ms. This configures the maximum duration of an entry in the stick-table since it was last created, refreshed or matched. The maximum duration is slightly above 24 days.",
				Computed:            true,
			},
			"stickiness_http_err_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request error rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"stickiness_http_req_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average HTTP request rate over that period, in requests per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"stickiness_length": schema.Int64Attribute{
				MarkdownDescription: "Specify the maximum length for a value in the stick-table.",
				Computed:            true,
			},
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "Choose the type of data that should be stored in this stick-table.",
				Computed:            true,
			},
			"stickiness_sess_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming session rate over that period, in sessions per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"stickiness_size": schema.StringAttribute{
				MarkdownDescription: "Enter a number followed by one of the supported suffixes k, m, g. This configures the maximum number of entries that can fit in the table. This value directly impacts memory usage. Count approximately 50 bytes per entry, plus the size of a string if any.",
				Computed:            true,
			},
			"tuning_max_connections": schema.Int64Attribute{
				MarkdownDescription: "Set the maximum number of concurrent connections for this Public Service.",
				Computed:            true,
			},
			"tuning_shards": schema.StringAttribute{
				MarkdownDescription: "This option automatically creates the specified number of listeners for every IP:port combination and evenly distributes them among available threads.",
				Computed:            true,
			},
			"tuning_timeout_client": schema.StringAttribute{
				MarkdownDescription: "Set the maximum inactivity time on the client side. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"tuning_timeout_http_keep_alive": schema.StringAttribute{
				MarkdownDescription: "Set the maximum allowed time to wait for a new HTTP request to appear. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"tuning_timeout_http_req": schema.StringAttribute{
				MarkdownDescription: "Set the maximum allowed time to wait for a complete HTTP request. In order to offer DoS protection, it may be required to lower the maximum accepted time to receive a complete HTTP request without affecting the client timeout. This helps protecting against established connections on which nothing is sent. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
		},
	}
}

func convertHaproxyFrontendSchemaToStruct(d *HaproxyFrontendResourceModel) (*haproxy.Frontend, error) {
	return &haproxy.Frontend{
		AdvertisedProtocols:          tools.SetToString(d.AdvertisedProtocols),
		BasicAuthEnabled:             tools.BoolToString(d.BasicAuthEnabled.ValueBool()),
		BasicAuthGroups:              tools.SetToString(d.BasicAuthGroups),
		BasicAuthUsers:               tools.SetToString(d.BasicAuthUsers),
		Bind:                         tools.SetToString(d.Bind),
		BindOptions:                  d.BindOptions.ValueString(),
		ConnectionBehaviour:          api.SelectedMap(d.ConnectionBehaviour.ValueString()),
		CustomOptions:                d.CustomOptions.ValueString(),
		DefaultBackend:               api.SelectedMap(d.DefaultBackend.ValueString()),
		Description:                  d.Description.ValueString(),
		Enabled:                      tools.BoolToString(d.Enabled.ValueBool()),
		ForwardFor:                   tools.BoolToString(d.ForwardFor.ValueBool()),
		Http2Enabled:                 tools.BoolToString(d.Http2Enabled.ValueBool()),
		Http2EnabledNontls:           tools.BoolToString(d.Http2EnabledNontls.ValueBool()),
		LinkedActions:                tools.SetToString(d.LinkedActions),
		LinkedCpuAffinityRules:       tools.SetToString(d.LinkedCpuAffinityRules),
		LinkedErrorfiles:             tools.SetToString(d.LinkedErrorfiles),
		LoggingDetailedLog:           tools.BoolToString(d.LoggingDetailedLog.ValueBool()),
		LoggingDontLogNormal:         tools.BoolToString(d.LoggingDontLogNormal.ValueBool()),
		LoggingDontLogNull:           tools.BoolToString(d.LoggingDontLogNull.ValueBool()),
		LoggingLogSeparateErrors:     tools.BoolToString(d.LoggingLogSeparateErrors.ValueBool()),
		LoggingSocketStats:           tools.BoolToString(d.LoggingSocketStats.ValueBool()),
		Mode:                         api.SelectedMap(d.Mode.ValueString()),
		Name:                         d.Name.ValueString(),
		PrometheusEnabled:            tools.BoolToString(d.PrometheusEnabled.ValueBool()),
		PrometheusPath:               d.PrometheusPath.ValueString(),
		SslAdvancedEnabled:           tools.BoolToString(d.SslAdvancedEnabled.ValueBool()),
		SslBindOptions:               tools.SetToString(d.SslBindOptions),
		SslCertificates:              tools.SetToString(d.SslCertificates),
		SslCipherList:                d.SslCipherList.ValueString(),
		SslCipherSuites:              d.SslCipherSuites.ValueString(),
		SslClientAuthCAs:             tools.SetToString(d.SslClientAuthCAs),
		SslClientAuthCRLs:            tools.SetToString(d.SslClientAuthCRLs),
		SslClientAuthEnabled:         tools.BoolToString(d.SslClientAuthEnabled.ValueBool()),
		SslClientAuthVerify:          api.SelectedMap(d.SslClientAuthVerify.ValueString()),
		SslCustomOptions:             d.SslCustomOptions.ValueString(),
		SslDefaultCertificate:        api.SelectedMap(d.SslDefaultCertificate.ValueString()),
		SslEnabled:                   tools.BoolToString(d.SslEnabled.ValueBool()),
		SslHstsEnabled:               tools.BoolToString(d.SslHstsEnabled.ValueBool()),
		SslHstsIncludeSubDomains:     tools.BoolToString(d.SslHstsIncludeSubDomains.ValueBool()),
		SslHstsMaxAge:                tools.Int64ToStringNegative(d.SslHstsMaxAge.ValueInt64()),
		SslHstsPreload:               tools.BoolToString(d.SslHstsPreload.ValueBool()),
		SslMaxVersion:                api.SelectedMap(d.SslMaxVersion.ValueString()),
		SslMinVersion:                api.SelectedMap(d.SslMinVersion.ValueString()),
		StickinessBytesInRatePeriod:  d.StickinessBytesInRatePeriod.ValueString(),
		StickinessBytesOutRatePeriod: d.StickinessBytesOutRatePeriod.ValueString(),
		StickinessConnRatePeriod:     d.StickinessConnRatePeriod.ValueString(),
		StickinessCounter:            tools.BoolToString(d.StickinessCounter.ValueBool()),
		StickinessCounterKey:         d.StickinessCounterKey.ValueString(),
		StickinessDataTypes:          tools.SetToString(d.StickinessDataTypes),
		StickinessExpire:             d.StickinessExpire.ValueString(),
		StickinessHttpErrRatePeriod:  d.StickinessHttpErrRatePeriod.ValueString(),
		StickinessHttpReqRatePeriod:  d.StickinessHttpReqRatePeriod.ValueString(),
		StickinessLength:             tools.Int64ToStringNegative(d.StickinessLength.ValueInt64()),
		StickinessPattern:            api.SelectedMap(d.StickinessPattern.ValueString()),
		StickinessSessRatePeriod:     d.StickinessSessRatePeriod.ValueString(),
		StickinessSize:               d.StickinessSize.ValueString(),
		TuningMaxConnections:         tools.Int64ToStringNegative(d.TuningMaxConnections.ValueInt64()),
		TuningShards:                 d.TuningShards.ValueString(),
		TuningTimeoutClient:          d.TuningTimeoutClient.ValueString(),
		TuningTimeoutHttpKeepAlive:   d.TuningTimeoutHttpKeepAlive.ValueString(),
		TuningTimeoutHttpReq:         d.TuningTimeoutHttpReq.ValueString(),
	}, nil
}

func convertHaproxyFrontendStructToSchema(d *haproxy.Frontend) (*HaproxyFrontendResourceModel, error) {
	return &HaproxyFrontendResourceModel{
		AdvertisedProtocols:          tools.StringToSet(d.AdvertisedProtocols),
		BasicAuthEnabled:             types.BoolValue(tools.StringToBool(d.BasicAuthEnabled)),
		BasicAuthGroups:              tools.StringToSet(d.BasicAuthGroups),
		BasicAuthUsers:               tools.StringToSet(d.BasicAuthUsers),
		Bind:                         tools.StringToSet(d.Bind),
		BindOptions:                  types.StringValue(d.BindOptions),
		ConnectionBehaviour:          types.StringValue(d.ConnectionBehaviour.String()),
		CustomOptions:                types.StringValue(d.CustomOptions),
		DefaultBackend:               types.StringValue(d.DefaultBackend.String()),
		Description:                  types.StringValue(d.Description),
		Enabled:                      types.BoolValue(tools.StringToBool(d.Enabled)),
		ForwardFor:                   types.BoolValue(tools.StringToBool(d.ForwardFor)),
		Http2Enabled:                 types.BoolValue(tools.StringToBool(d.Http2Enabled)),
		Http2EnabledNontls:           types.BoolValue(tools.StringToBool(d.Http2EnabledNontls)),
		LinkedActions:                tools.StringToSet(d.LinkedActions),
		LinkedCpuAffinityRules:       tools.StringToSet(d.LinkedCpuAffinityRules),
		LinkedErrorfiles:             tools.StringToSet(d.LinkedErrorfiles),
		LoggingDetailedLog:           types.BoolValue(tools.StringToBool(d.LoggingDetailedLog)),
		LoggingDontLogNormal:         types.BoolValue(tools.StringToBool(d.LoggingDontLogNormal)),
		LoggingDontLogNull:           types.BoolValue(tools.StringToBool(d.LoggingDontLogNull)),
		LoggingLogSeparateErrors:     types.BoolValue(tools.StringToBool(d.LoggingLogSeparateErrors)),
		LoggingSocketStats:           types.BoolValue(tools.StringToBool(d.LoggingSocketStats)),
		Mode:                         types.StringValue(d.Mode.String()),
		Name:                         types.StringValue(d.Name),
		PrometheusEnabled:            types.BoolValue(tools.StringToBool(d.PrometheusEnabled)),
		PrometheusPath:               types.StringValue(d.PrometheusPath),
		SslAdvancedEnabled:           types.BoolValue(tools.StringToBool(d.SslAdvancedEnabled)),
		SslBindOptions:               tools.StringToSet(d.SslBindOptions),
		SslCertificates:              tools.StringToSet(d.SslCertificates),
		SslCipherList:                types.StringValue(d.SslCipherList),
		SslCipherSuites:              types.StringValue(d.SslCipherSuites),
		SslClientAuthCAs:             tools.StringToSet(d.SslClientAuthCAs),
		SslClientAuthCRLs:            tools.StringToSet(d.SslClientAuthCRLs),
		SslClientAuthEnabled:         types.BoolValue(tools.StringToBool(d.SslClientAuthEnabled)),
		SslClientAuthVerify:          types.StringValue(d.SslClientAuthVerify.String()),
		SslCustomOptions:             types.StringValue(d.SslCustomOptions),
		SslDefaultCertificate:        types.StringValue(d.SslDefaultCertificate.String()),
		SslEnabled:                   types.BoolValue(tools.StringToBool(d.SslEnabled)),
		SslHstsEnabled:               types.BoolValue(tools.StringToBool(d.SslHstsEnabled)),
		SslHstsIncludeSubDomains:     types.BoolValue(tools.StringToBool(d.SslHstsIncludeSubDomains)),
		SslHstsMaxAge:                types.Int64Value(tools.StringToInt64(d.SslHstsMaxAge)),
		SslHstsPreload:               types.BoolValue(tools.StringToBool(d.SslHstsPreload)),
		SslMaxVersion:                types.StringValue(d.SslMaxVersion.String()),
		SslMinVersion:                types.StringValue(d.SslMinVersion.String()),
		StickinessBytesInRatePeriod:  types.StringValue(d.StickinessBytesInRatePeriod),
		StickinessBytesOutRatePeriod: types.StringValue(d.StickinessBytesOutRatePeriod),
		StickinessConnRatePeriod:     types.StringValue(d.StickinessConnRatePeriod),
		StickinessCounter:            types.BoolValue(tools.StringToBool(d.StickinessCounter)),
		StickinessCounterKey:         types.StringValue(d.StickinessCounterKey),
		StickinessDataTypes:          tools.StringToSet(d.StickinessDataTypes),
		StickinessExpire:             types.StringValue(d.StickinessExpire),
		StickinessHttpErrRatePeriod:  types.StringValue(d.StickinessHttpErrRatePeriod),
		StickinessHttpReqRatePeriod:  types.StringValue(d.StickinessHttpReqRatePeriod),
		StickinessLength:             types.Int64Value(tools.StringToInt64(d.StickinessLength)),
		StickinessPattern:            types.StringValue(d.StickinessPattern.String()),
		StickinessSessRatePeriod:     types.StringValue(d.StickinessSessRatePeriod),
		StickinessSize:               types.StringValue(d.StickinessSize),
		TuningMaxConnections:         types.Int64Value(tools.StringToInt64(d.TuningMaxConnections)),
		TuningShards:                 types.StringValue(d.TuningShards),
		TuningTimeoutClient:          types.StringValue(d.TuningTimeoutClient),
		TuningTimeoutHttpKeepAlive:   types.StringValue(d.TuningTimeoutHttpKeepAlive),
		TuningTimeoutHttpReq:         types.StringValue(d.TuningTimeoutHttpReq),
	}, nil
}
