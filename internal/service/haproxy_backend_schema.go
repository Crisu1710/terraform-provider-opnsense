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

// HaproxyBackendResourceModel describes the resource data model.
type HaproxyBackendResourceModel struct {
	Algorithm                    types.String `tfsdk:"algorithm"`
	BaAdvertisedProtocols        types.Set    `tfsdk:"ba_advertised_protocols"`
	BasicAuthEnabled             types.Bool   `tfsdk:"basic_auth_enabled"`
	BasicAuthGroups              types.Set    `tfsdk:"basic_auth_groups"`
	BasicAuthUsers               types.Set    `tfsdk:"basic_auth_users"`
	CheckDownInterval            types.Int64  `tfsdk:"check_down_interval"`
	CheckInterval                types.Int64  `tfsdk:"check_interval"`
	CustomOptions                types.String `tfsdk:"custom_options"`
	Description                  types.String `tfsdk:"description"`
	Enabled                      types.Bool   `tfsdk:"enabled"`
	HealthCheck                  types.String `tfsdk:"health_check"`
	HealthCheckEnabled           types.Bool   `tfsdk:"health_check_enabled"`
	HealthCheckFall              types.Int64  `tfsdk:"health_check_fall"`
	HealthCheckLogStatus         types.Bool   `tfsdk:"health_check_log_status"`
	HealthCheckRise              types.Int64  `tfsdk:"health_check_rise"`
	Http2Enabled                 types.Bool   `tfsdk:"http2enabled"`
	Http2EnabledNontls           types.Bool   `tfsdk:"http2enabled_nontls"`
	LinkedActions                types.Set    `tfsdk:"linked_actions"`
	LinkedErrorfiles             types.Set    `tfsdk:"linked_errorfiles"`
	LinkedFcgi                   types.String `tfsdk:"linked_fcgi"`
	LinkedMailer                 types.String `tfsdk:"linked_mailer"`
	LinkedResolver               types.String `tfsdk:"linked_resolver"`
	LinkedServers                types.Set    `tfsdk:"linked_servers"`
	Mode                         types.String `tfsdk:"mode"`
	Name                         types.String `tfsdk:"name"`
	Persistence                  types.String `tfsdk:"persistence"`
	PersistenceCookiemode        types.String `tfsdk:"persistence_cookiemode"`
	PersistenceCookiename        types.String `tfsdk:"persistence_cookiename"`
	PersistenceStripquotes       types.Bool   `tfsdk:"persistence_stripquotes"`
	ProxyProtocol                types.String `tfsdk:"proxy_protocol"`
	RandomDraws                  types.Int64  `tfsdk:"random_draws"`
	ResolvePrefer                types.String `tfsdk:"resolve_prefer"`
	ResolverOpts                 types.Set    `tfsdk:"resolver_opts"`
	Source                       types.String `tfsdk:"source"`
	StickinessBytesInRatePeriod  types.String `tfsdk:"stickiness_bytes_in_rate_period"`
	StickinessBytesOutRatePeriod types.String `tfsdk:"stickiness_bytes_out_rate_period"`
	StickinessConnRatePeriod     types.String `tfsdk:"stickiness_conn_rate_period"`
	StickinessCookielength       types.Int64  `tfsdk:"stickiness_cookielength"`
	StickinessCookiename         types.String `tfsdk:"stickiness_cookiename"`
	StickinessDataTypes          types.Set    `tfsdk:"stickiness_data_types"`
	StickinessExpire             types.String `tfsdk:"stickiness_expire"`
	StickinessHttpErrRatePeriod  types.String `tfsdk:"stickiness_http_err_rate_period"`
	StickinessHttpReqRatePeriod  types.String `tfsdk:"stickiness_http_req_rate_period"`
	StickinessPattern            types.String `tfsdk:"stickiness_pattern"`
	StickinessSessRatePeriod     types.String `tfsdk:"stickiness_sess_rate_period"`
	StickinessSize               types.String `tfsdk:"stickiness_size"`
	TuningCaching                types.Bool   `tfsdk:"tuning_caching"`
	TuningDefaultserver          types.String `tfsdk:"tuning_defaultserver"`
	TuningHttpreuse              types.String `tfsdk:"tuning_httpreuse"`
	TuningNoport                 types.Bool   `tfsdk:"tuning_noport"`
	TuningRetries                types.Int64  `tfsdk:"tuning_retries"`
	TuningTimeoutCheck           types.String `tfsdk:"tuning_timeout_check"`
	TuningTimeoutConnect         types.String `tfsdk:"tuning_timeout_connect"`
	TuningTimeoutServer          types.String `tfsdk:"tuning_timeout_server"`

	Id types.String `tfsdk:"id"`
}

func haproxyBackendResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Host backends can be used to change DNS results from client queries or to add custom DNS records.",

		Attributes: map[string]schema.Attribute{
			"algorithm": schema.StringAttribute{
				MarkdownDescription: "Define the load balancing algorithm to be used in a Backend Pool. See the [HAProxy documentation](https://docs.haproxy.org/2.6/configuration.html#balance) for a full description.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("source", "roundrobin", "static-rr", "leastconn", "uri", "random"),
				},
				Default: stringdefault.StaticString("source"),
			},
			"ba_advertised_protocols": schema.SetAttribute{
				MarkdownDescription: "When using the TLS ALPN extension, HAProxy advertises the specified protocol list as supported on top of ALPN. TLS must be enabled.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(stringvalidator.OneOf("h2", "http11", "http10")),
				},
				Default: setdefault.StaticValue(tools.AppendSetValue([]string{"h2", "http11"})),
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP Basic Authentication.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"basic_auth_groups": schema.SetAttribute{
				MarkdownDescription: "list of groups seperated by ,",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"basic_auth_users": schema.SetAttribute{
				MarkdownDescription: "list of users seperated by ,",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"check_down_interval": schema.Int64Attribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on a configured server when the server state is DOWN. If it is not set HAProxy uses the check interval.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"check_interval": schema.Int64Attribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on all configured servers. This setting takes precedence over default values in health monitors and real servers.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "These line will be added to the HAProxy backend configuration.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this Backend Pool.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this Backend Pool",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"health_check": schema.StringAttribute{
				MarkdownDescription: "Select Health Monitor for servers in this backend.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_check_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable health checking.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"health_check_fall": schema.Int64Attribute{
				MarkdownDescription: "The number of consecutive unsuccessful health checks before a server is considered as unavailable.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"health_check_log_status": schema.BoolAttribute{
				MarkdownDescription: "Enable to log health check status updates.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"health_check_rise": schema.Int64Attribute{
				MarkdownDescription: "The number of consecutive successful health checks before a server is considered as available.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"http2enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable support for end-to-end HTTP/2 communication.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http2enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "Enable support for HTTP/2 even if TLS is not enabled.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"linked_actions": schema.SetAttribute{
				MarkdownDescription: "list of rules seperated by ```,``` to be included in this Backend Pool.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"linked_errorfiles": schema.SetAttribute{
				MarkdownDescription: "list of error messages seperated by ```,``` to be included in this Backend Pool.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"linked_fcgi": schema.StringAttribute{
				MarkdownDescription: "The FastCGI application that should be used for all servers in this backend.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_mailer": schema.StringAttribute{
				MarkdownDescription: "Set an e-mail alert configuration. An e-mail is sent when the state of a server changes.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "The the custom resolver configuration that should be used for all servers in this backend.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_servers": schema.SetAttribute{
				MarkdownDescription: "Link the server(s) to this backend.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Set the running mode or protocol of the Backend Pool. Usually the Public Service and the Backend Pool are in the same mode.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http", "tcp"),
				},
				Default: stringdefault.StaticString("http"),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this Backend Pool.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"persistence": schema.StringAttribute{
				MarkdownDescription: "Choose (sticktable, cookie) how HAProxy should track user-to-server mappings.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "sticktable", "cookie"),
				},
				Default: stringdefault.StaticString("sticktable"),
			},
			"persistence_cookiemode": schema.StringAttribute{
				MarkdownDescription: "(piggyback or new) Usually it is better to reuse an existing cookie. In this case HAProxy prefixes the cookie with the required information. See the [HAProxy documentation](https://docs.haproxy.org/2.6/configuration.html#4.2-cookie) for a full description.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("piggyback", "new"),
				},
				Default: stringdefault.StaticString("piggyback"),
			},
			"persistence_cookiename": schema.StringAttribute{
				MarkdownDescription: "Cookie name to use for persistence.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("SRVCOOKIE"),
			},
			"persistence_stripquotes": schema.BoolAttribute{
				MarkdownDescription: "Enable to automatically strip quotes from the cookie value.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"proxy_protocol": schema.StringAttribute{
				MarkdownDescription: "Enforces use of the PROXY protocol over any connection established to the configured servers. (v1, v2, unset to deactivate)",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "v1", "v2"),
				},
				Default: stringdefault.StaticString(""),
			},
			"random_draws": schema.Int64Attribute{
				MarkdownDescription: "When using the Random Balancing Algorithm, this value indicates the number of draws before selecting the least loaded of these servers.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2),
			},
			"resolve_prefer": schema.StringAttribute{
				MarkdownDescription: "When DNS resolution is enabled for a server and multiple IP addresses from different families are returned, HAProxy will prefer using an IP address from the selected family. (ipv4, ipv6)",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "ipv4", "ipv6"),
				},
				Default: stringdefault.StaticString(""),
			},
			"resolver_opts": schema.SetAttribute{
				MarkdownDescription: "Add resolver options seperated by ```,``` (allow-dup-ip, allow-dup-ip, prevent-dup-ip).",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(stringvalidator.OneOf("", "allow-dup-ip", "ignore-weight", "prevent-dup-ip")),
				},
				Default: setdefault.StaticValue(tools.EmptySetValue()),
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Sets the source address which will be used when connecting to the server(s).",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
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
			"stickiness_cookielength": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of characters that will be stored in the stick table (if appropiate table type is selected).",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"stickiness_cookiename": schema.StringAttribute{
				MarkdownDescription: "Cookie name to use for stick table (if appropiate table type is selected).",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"stickiness_data_types": schema.SetAttribute{
				MarkdownDescription: "This is used to store additional information in the stick-table. ",
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
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "The request pattern to associate a user to a server (sourceipv4, sourceipv6, cookievalue, rdpcookie). See the [HAProxy documentation](https://docs.haproxy.org/2.6/configuration.html#stick%20on) for a full description.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "sourceipv4", "sourceipv6", "cookievalue", "rdpcookie"),
				},
				Default: stringdefault.StaticString("sourceipv4"),
			},
			"stickiness_sess_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming session rate over that period, in sessions per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_size": schema.StringAttribute{
				MarkdownDescription: "Enter a number followed by one of the supported suffixes k, m, g.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("50k"),
			},
			"tuning_caching": schema.BoolAttribute{
				MarkdownDescription: "Enable caching of responses from this backend. The HAProxy cache must be enabled under Settings before this will have any effect.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tuning_defaultserver": schema.StringAttribute{
				MarkdownDescription: "Default option for all server entries.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_httpreuse": schema.StringAttribute{
				MarkdownDescription: "Declare how idle HTTP connections may be shared between requests. (never, safe, aggressive, always)",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "never", "safe", "aggressive", "always"),
				},
				Default: stringdefault.StaticString("safe"),
			},
			"tuning_noport": schema.BoolAttribute{
				MarkdownDescription: "Don't use port on server, use the same port as frontend receive. If enabled, require port check in server.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tuning_retries": schema.Int64Attribute{
				MarkdownDescription: "Set the number of retries to perform on a server after a connection failure.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"tuning_timeout_check": schema.StringAttribute{
				MarkdownDescription: "Sets an additional read timeout for running health checks on a server. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_connect": schema.StringAttribute{
				MarkdownDescription: "Set the maximum time to wait for a connection attempt to a server to succeed. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_server": schema.StringAttribute{
				MarkdownDescription: "Set the maximum inactivity time on the server side. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the host backend.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func HaproxyBackendDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Host backends can be used to change DNS results from client queries or to add custom DNS records.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"algorithm": schema.StringAttribute{
				MarkdownDescription: "Define the load balancing algorithm to be used in a Backend Pool. See the [HAProxy documentation](https://docs.haproxy.org/2.6/configuration.html#balance) for a full description.",
				Computed:            true,
			},
			"ba_advertised_protocols": schema.SetAttribute{
				MarkdownDescription: "When using the TLS ALPN extension, HAProxy advertises the specified protocol list as supported on top of ALPN. TLS must be enabled.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable HTTP Basic Authentication.",
				Computed:            true,
			},
			"basic_auth_groups": schema.SetAttribute{
				MarkdownDescription: "list of groups seperated by ,",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"basic_auth_users": schema.SetAttribute{
				MarkdownDescription: "list of users seperated by ,",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"check_down_interval": schema.Int64Attribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on a configured server when the server state is DOWN. If it is not set HAProxy uses the check interval.",
				Computed:            true,
			},
			"check_interval": schema.Int64Attribute{
				MarkdownDescription: "Sets the interval (in milliseconds) for running health checks on all configured servers. This setting takes precedence over default values in health monitors and real servers.",
				Computed:            true,
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "These line will be added to the HAProxy backend configuration.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this Backend Pool.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this Backend Pool",
				Computed:            true,
			},
			"health_check": schema.StringAttribute{
				MarkdownDescription: "Select Health Monitor for servers in this backend.",
				Computed:            true,
			},
			"health_check_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable health checking.",
				Computed:            true,
			},
			"health_check_fall": schema.Int64Attribute{
				MarkdownDescription: "The number of consecutive unsuccessful health checks before a server is considered as unavailable.",
				Computed:            true,
			},
			"health_check_log_status": schema.BoolAttribute{
				MarkdownDescription: "Enable to log health check status updates.",
				Computed:            true,
			},
			"health_check_rise": schema.Int64Attribute{
				MarkdownDescription: "The number of consecutive successful health checks before a server is considered as available.",
				Computed:            true,
			},
			"http2enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable support for end-to-end HTTP/2 communication.",
				Computed:            true,
			},
			"http2enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "Enable support for HTTP/2 even if TLS is not enabled.",
				Computed:            true,
			},
			"linked_actions": schema.SetAttribute{
				MarkdownDescription: "list of rules seperated by ```,``` to be included in this Backend Pool.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"linked_errorfiles": schema.SetAttribute{
				MarkdownDescription: "list of error messages seperated by ```,``` to be included in this Backend Pool.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"linked_fcgi": schema.StringAttribute{
				MarkdownDescription: "The FastCGI application that should be used for all servers in this backend.",
				Computed:            true,
			},
			"linked_mailer": schema.StringAttribute{
				MarkdownDescription: "Set an e-mail alert configuration. An e-mail is sent when the state of a server changes.",
				Computed:            true,
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "The the custom resolver configuration that should be used for all servers in this backend.",
				Computed:            true,
			},
			"linked_servers": schema.SetAttribute{
				MarkdownDescription: "Link the server(s) to this backend.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Set the running mode or protocol of the Backend Pool. Usually the Public Service and the Backend Pool are in the same mode.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name to identify this Backend Pool.",
				Required:            true,
			},
			"persistence": schema.StringAttribute{
				MarkdownDescription: "Choose (sticktable, cookie) how HAProxy should track user-to-server mappings.",
				Computed:            true,
			},
			"persistence_cookiemode": schema.StringAttribute{
				MarkdownDescription: "(piggyback or new) Usually it is better to reuse an existing cookie. In this case HAProxy prefixes the cookie with the required information. See the [HAProxy documentation](https://docs.haproxy.org/2.6/configuration.html#4.2-cookie) for a full description.",
				Computed:            true,
			},
			"persistence_cookiename": schema.StringAttribute{
				MarkdownDescription: "Cookie name to use for persistence.",
				Computed:            true,
			},
			"persistence_stripquotes": schema.BoolAttribute{
				MarkdownDescription: "Enable to automatically strip quotes from the cookie value.",
				Computed:            true,
			},
			"proxy_protocol": schema.StringAttribute{
				MarkdownDescription: "Enforces use of the PROXY protocol over any connection established to the configured servers. (v1, v2, unset to deactivate)",
				Computed:            true,
			},
			"random_draws": schema.Int64Attribute{
				MarkdownDescription: "When using the Random Balancing Algorithm, this value indicates the number of draws before selecting the least loaded of these servers.",
				Computed:            true,
			},
			"resolve_prefer": schema.StringAttribute{
				MarkdownDescription: "When DNS resolution is enabled for a server and multiple IP addresses from different families are returned, HAProxy will prefer using an IP address from the selected family. (ipv4, ipv6)",
				Computed:            true,
			},
			"resolver_opts": schema.SetAttribute{
				MarkdownDescription: "Add resolver options seperated by ```,``` (allow-dup-ip, allow-dup-ip, prevent-dup-ip).",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"source": schema.StringAttribute{
				MarkdownDescription: "Sets the source address which will be used when connecting to the server(s).",
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
			"stickiness_cookielength": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of characters that will be stored in the stick table (if appropiate table type is selected).",
				Computed:            true,
			},
			"stickiness_cookiename": schema.StringAttribute{
				MarkdownDescription: "Cookie name to use for stick table (if appropiate table type is selected).",
				Computed:            true,
			},
			"stickiness_data_types": schema.SetAttribute{
				MarkdownDescription: "This is used to store additional information in the stick-table. ",
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
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "The request pattern to associate a user to a server (sourceipv4, sourceipv6, cookievalue, rdpcookie). See the [HAProxy documentation](https://docs.haproxy.org/2.6/configuration.html#stick%20on) for a full description.",
				Computed:            true,
			},
			"stickiness_sess_rate_period": schema.StringAttribute{
				MarkdownDescription: "The length of the period over which the average is measured. It reports the average incoming session rate over that period, in sessions per period. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"stickiness_size": schema.StringAttribute{
				MarkdownDescription: "Enter a number followed by one of the supported suffixes k, m, g.",
				Computed:            true,
			},
			"tuning_caching": schema.BoolAttribute{
				MarkdownDescription: "Enable caching of responses from this backend. The HAProxy cache must be enabled under Settings before this will have any effect.",
				Computed:            true,
			},
			"tuning_defaultserver": schema.StringAttribute{
				MarkdownDescription: "Default option for all server entries.",
				Computed:            true,
			},
			"tuning_httpreuse": schema.StringAttribute{
				MarkdownDescription: "Declare how idle HTTP connections may be shared between requests. (never, safe, aggressive, always)",
				Computed:            true,
			},
			"tuning_noport": schema.BoolAttribute{
				MarkdownDescription: "Don't use port on server, use the same port as frontend receive. If enabled, require port check in server.",
				Computed:            true,
			},
			"tuning_retries": schema.Int64Attribute{
				MarkdownDescription: "Set the number of retries to perform on a server after a connection failure.",
				Computed:            true,
			},
			"tuning_timeout_check": schema.StringAttribute{
				MarkdownDescription: "Sets an additional read timeout for running health checks on a server. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"tuning_timeout_connect": schema.StringAttribute{
				MarkdownDescription: "Set the maximum time to wait for a connection attempt to a server to succeed. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
			"tuning_timeout_server": schema.StringAttribute{
				MarkdownDescription: "Set the maximum inactivity time on the server side. Defaults to milliseconds. Optionally the unit may be specified as either d, h, m, s, ms or us.",
				Computed:            true,
			},
		},
	}
}

func convertHaproxyBackendSchemaToStruct(d *HaproxyBackendResourceModel) (*haproxy.Backend, error) {
	return &haproxy.Backend{
		Algorithm:                    api.SelectedMap(d.Algorithm.ValueString()),
		BaAdvertisedProtocols:        tools.SetToString(d.BaAdvertisedProtocols),
		BasicAuthEnabled:             tools.BoolToString(d.BasicAuthEnabled.ValueBool()),
		BasicAuthGroups:              tools.SetToString(d.BasicAuthGroups),
		BasicAuthUsers:               tools.SetToString(d.BasicAuthUsers),
		CheckDownInterval:            tools.Int64ToStringNegative(d.CheckDownInterval.ValueInt64()),
		CheckInterval:                tools.Int64ToStringNegative(d.CheckInterval.ValueInt64()),
		CustomOptions:                d.CustomOptions.ValueString(),
		Description:                  d.Description.ValueString(),
		Enabled:                      tools.BoolToString(d.Enabled.ValueBool()),
		HealthCheck:                  api.SelectedMap(d.HealthCheck.ValueString()),
		HealthCheckEnabled:           tools.BoolToString(d.HealthCheckEnabled.ValueBool()),
		HealthCheckFall:              tools.Int64ToStringNegative(d.HealthCheckFall.ValueInt64()),
		HealthCheckLogStatus:         tools.BoolToString(d.HealthCheckLogStatus.ValueBool()),
		HealthCheckRise:              tools.Int64ToStringNegative(d.HealthCheckRise.ValueInt64()),
		Http2Enabled:                 tools.BoolToString(d.Http2Enabled.ValueBool()),
		Http2EnabledNontls:           tools.BoolToString(d.Http2EnabledNontls.ValueBool()),
		LinkedActions:                tools.SetToString(d.LinkedActions),
		LinkedErrorfiles:             tools.SetToString(d.LinkedErrorfiles),
		LinkedFcgi:                   api.SelectedMap(d.LinkedFcgi.ValueString()),
		LinkedMailer:                 api.SelectedMap(d.LinkedMailer.ValueString()),
		LinkedResolver:               api.SelectedMap(d.LinkedResolver.ValueString()),
		LinkedServers:                tools.SetToString(d.LinkedServers),
		Mode:                         api.SelectedMap(d.Mode.ValueString()),
		Name:                         d.Name.ValueString(),
		Persistence:                  api.SelectedMap(d.Persistence.ValueString()),
		PersistenceCookiemode:        api.SelectedMap(d.PersistenceCookiemode.ValueString()),
		PersistenceCookiename:        d.PersistenceCookiename.ValueString(),
		PersistenceStripquotes:       tools.BoolToString(d.PersistenceStripquotes.ValueBool()),
		ProxyProtocol:                api.SelectedMap(d.ProxyProtocol.ValueString()),
		RandomDraws:                  tools.Int64ToStringNegative(d.RandomDraws.ValueInt64()),
		ResolvePrefer:                api.SelectedMap(d.ResolvePrefer.ValueString()),
		ResolverOpts:                 tools.SetToString(d.ResolverOpts),
		Source:                       d.Source.ValueString(),
		StickinessBytesInRatePeriod:  d.StickinessBytesInRatePeriod.ValueString(),
		StickinessBytesOutRatePeriod: d.StickinessBytesOutRatePeriod.ValueString(),
		StickinessConnRatePeriod:     d.StickinessConnRatePeriod.ValueString(),
		StickinessCookielength:       tools.Int64ToStringNegative(d.StickinessCookielength.ValueInt64()),
		StickinessCookiename:         d.StickinessCookiename.ValueString(),
		StickinessDataTypes:          tools.SetToString(d.StickinessDataTypes),
		StickinessExpire:             d.StickinessExpire.ValueString(),
		StickinessHttpErrRatePeriod:  d.StickinessHttpErrRatePeriod.ValueString(),
		StickinessHttpReqRatePeriod:  d.StickinessHttpReqRatePeriod.ValueString(),
		StickinessPattern:            api.SelectedMap(d.StickinessPattern.ValueString()),
		StickinessSessRatePeriod:     d.StickinessSessRatePeriod.ValueString(),
		StickinessSize:               d.StickinessSize.ValueString(),
		TuningCaching:                tools.BoolToString(d.TuningCaching.ValueBool()),
		TuningDefaultserver:          d.TuningDefaultserver.ValueString(),
		TuningHttpreuse:              api.SelectedMap(d.TuningHttpreuse.ValueString()),
		TuningNoport:                 tools.BoolToString(d.TuningNoport.ValueBool()),
		TuningRetries:                tools.Int64ToStringNegative(d.TuningRetries.ValueInt64()),
		TuningTimeoutCheck:           d.TuningTimeoutCheck.ValueString(),
		TuningTimeoutConnect:         d.TuningTimeoutConnect.ValueString(),
		TuningTimeoutServer:          d.TuningTimeoutServer.ValueString(),
	}, nil
}

func convertHaproxyBackendStructToSchema(d *haproxy.Backend) (*HaproxyBackendResourceModel, error) {
	return &HaproxyBackendResourceModel{
		Algorithm:                    types.StringValue(d.Algorithm.String()),
		BaAdvertisedProtocols:        tools.StringToSet(d.BaAdvertisedProtocols),
		BasicAuthEnabled:             types.BoolValue(tools.StringToBool(d.BasicAuthEnabled)),
		BasicAuthGroups:              tools.StringToSet(d.BasicAuthGroups),
		BasicAuthUsers:               tools.StringToSet(d.BasicAuthUsers),
		CheckDownInterval:            types.Int64Value(tools.StringToInt64(d.CheckDownInterval)),
		CheckInterval:                types.Int64Value(tools.StringToInt64(d.CheckInterval)),
		CustomOptions:                types.StringValue(d.CustomOptions),
		Description:                  types.StringValue(d.Description),
		Enabled:                      types.BoolValue(tools.StringToBool(d.Enabled)),
		HealthCheck:                  types.StringValue(d.HealthCheck.String()),
		HealthCheckEnabled:           types.BoolValue(tools.StringToBool(d.HealthCheckEnabled)),
		HealthCheckFall:              types.Int64Value(tools.StringToInt64(d.HealthCheckFall)),
		HealthCheckLogStatus:         types.BoolValue(tools.StringToBool(d.HealthCheckLogStatus)),
		HealthCheckRise:              types.Int64Value(tools.StringToInt64(d.HealthCheckRise)),
		Http2Enabled:                 types.BoolValue(tools.StringToBool(d.Http2Enabled)),
		Http2EnabledNontls:           types.BoolValue(tools.StringToBool(d.Http2EnabledNontls)),
		LinkedActions:                tools.StringToSet(d.LinkedActions),
		LinkedErrorfiles:             tools.StringToSet(d.LinkedErrorfiles),
		LinkedFcgi:                   types.StringValue(d.LinkedFcgi.String()),
		LinkedMailer:                 types.StringValue(d.LinkedMailer.String()),
		LinkedResolver:               types.StringValue(d.LinkedResolver.String()),
		LinkedServers:                tools.StringToSet(d.LinkedServers),
		Mode:                         types.StringValue(d.Mode.String()),
		Name:                         types.StringValue(d.Name),
		Persistence:                  types.StringValue(d.Persistence.String()),
		PersistenceCookiemode:        types.StringValue(d.PersistenceCookiemode.String()),
		PersistenceCookiename:        types.StringValue(d.PersistenceCookiename),
		PersistenceStripquotes:       types.BoolValue(tools.StringToBool(d.PersistenceStripquotes)),
		ProxyProtocol:                types.StringValue(d.ProxyProtocol.String()),
		RandomDraws:                  types.Int64Value(tools.StringToInt64(d.RandomDraws)),
		ResolvePrefer:                types.StringValue(d.ResolvePrefer.String()),
		ResolverOpts:                 tools.StringToSet(d.ResolverOpts),
		Source:                       types.StringValue(d.Source),
		StickinessBytesInRatePeriod:  types.StringValue(d.StickinessBytesInRatePeriod),
		StickinessBytesOutRatePeriod: types.StringValue(d.StickinessBytesOutRatePeriod),
		StickinessConnRatePeriod:     types.StringValue(d.StickinessConnRatePeriod),
		StickinessCookielength:       types.Int64Value(tools.StringToInt64(d.StickinessCookielength)),
		StickinessCookiename:         types.StringValue(d.StickinessCookiename),
		StickinessDataTypes:          tools.StringToSet(d.StickinessDataTypes),
		StickinessExpire:             types.StringValue(d.StickinessExpire),
		StickinessHttpErrRatePeriod:  types.StringValue(d.StickinessHttpErrRatePeriod),
		StickinessHttpReqRatePeriod:  types.StringValue(d.StickinessHttpReqRatePeriod),
		StickinessPattern:            types.StringValue(d.StickinessPattern.String()),
		StickinessSessRatePeriod:     types.StringValue(d.StickinessHttpReqRatePeriod),
		StickinessSize:               types.StringValue(d.StickinessSize),
		TuningCaching:                types.BoolValue(tools.StringToBool(d.TuningCaching)),
		TuningDefaultserver:          types.StringValue(d.TuningDefaultserver),
		TuningHttpreuse:              types.StringValue(d.TuningHttpreuse.String()),
		TuningNoport:                 types.BoolValue(tools.StringToBool(d.TuningNoport)),
		TuningRetries:                types.Int64Value(tools.StringToInt64(d.TuningRetries)),
		TuningTimeoutCheck:           types.StringValue(d.TuningTimeoutCheck),
		TuningTimeoutConnect:         types.StringValue(d.TuningTimeoutConnect),
		TuningTimeoutServer:          types.StringValue(d.TuningTimeoutServer),
	}, nil
}
