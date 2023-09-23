package service

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/haproxy"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-opnsense/internal/tools"
)

// HaproxyFrontendResourceModel describes the resource data model.
type HaproxyFrontendResourceModel struct {
	AdvertisedProtocols          types.String `tfsdk:"advertised_protocols"`
	BasicAuthEnabled             types.Bool   `tfsdk:"basic_auth_enabled"`
	BasicAuthGroups              types.Set    `tfsdk:"basic_auth_groups"`
	BasicAuthUsers               types.Set    `tfsdk:"basic_auth_users"`
	Bind                         types.String `tfsdk:"bind"`
	BindOptions                  types.String `tfsdk:"bind_options"`
	ConnectionBehaviour          types.String `tfsdk:"connection_behaviour"`
	CustomOptions                types.String `tfsdk:"custom_options"`
	DefaultBackend               types.String `tfsdk:"default_backend"`
	Description                  types.String `tfsdk:"description"`
	Enabled                      types.Bool   `tfsdk:"enabled"`
	ForwardFor                   types.Bool   `tfsdk:"forward_for"`
	Http2Enabled                 types.Bool   `tfsdk:"http2enabled"`
	Http2EnabledNontls           types.Bool   `tfsdk:"http2enabled_nontls"`
	LinkedActions                types.String `tfsdk:"linked_actions"`
	LinkedCpuAffinityRules       types.String `tfsdk:"linked_cpu_affinity_rules"`
	LinkedErrorfiles             types.String `tfsdk:"linked_errorfiles"`
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
	SslCertificates              types.String `tfsdk:"ssl_certificates"`
	SslCipherList                types.String `tfsdk:"ssl_cipher_list"`
	SslCipherSuites              types.String `tfsdk:"ssl_cipher_suites"`
	SslClientAuthCAs             types.String `tfsdk:"ssl_client_auth_cas"`
	SslClientAuthCRLs            types.String `tfsdk:"ssl_client_auth_crls"`
	SslClientAuthEnabled         types.Bool   `tfsdk:"ssl_client_auth_enabled"`
	SslClientAuthVerify          types.String `tfsdk:"ssl_client_auth_verify"`
	SslCustomOptions             types.String `tfsdk:"ssl_custom_options"`
	SslDefaultCertificate        types.String `tfsdk:"ssl_default_certificate"`
	SslEnabled                   types.Bool   `tfsdk:"ssl_enabled"`
	SslHstsEnabled               types.Bool   `tfsdk:"ssl_hsts_enabled"`
	SslHstsIncludeSubDomains     types.Bool   `tfsdk:"ssl_hsts_include_sub_domains"`
	SslHstsMaxAge                types.Int64  `tfsdk:"ssl_hsts_max_age"`
	SslHstsPreload               types.String `tfsdk:"ssl_hsts_preload"`
	SslMaxVersion                types.String `tfsdk:"ssl_max_version"`
	SslMinVersion                types.String `tfsdk:"ssl_min_version"`
	StickinessBytesInRatePeriod  types.String `tfsdk:"stickiness_bytes_in_rate_period"`
	StickinessBytesOutRatePeriod types.String `tfsdk:"stickiness_bytes_out_rate_period"`
	StickinessConnRatePeriod     types.String `tfsdk:"stickiness_conn_rate_period"`
	StickinessCounter            types.Int64  `tfsdk:"stickiness_counter"`
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
			"advertised_protocols": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("h2,http11"),
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "ba_advertised_protocols",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"basic_auth_groups": schema.SetAttribute{
				MarkdownDescription: "basicAuthEnabled",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"basic_auth_users": schema.SetAttribute{
				MarkdownDescription: "basicAuthGroups",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"bind": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Required:            true,
			},
			"bind_options": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"connection_behaviour": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http-keep-alive"),
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"default_backend": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
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
			"forward_for": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http2enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http2enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"linked_actions": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_cpu_affinity_rules": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_errorfiles": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"logging_detailed_log": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_dont_log_normal": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_dont_log_null": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_log_separate_errors": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"logging_socket_stats": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http"),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Required:            true,
			},
			"prometheus_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"prometheus_path": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_advanced_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_bind_options": schema.SetAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"ssl_certificates": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_cipher_list": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256"),
			},
			"ssl_cipher_suites": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("TLS_AES_128_GCM_SHA256:TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256"),
			},
			"ssl_client_auth_cas": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_client_auth_crls": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_client_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_client_auth_verify": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("required"),
			},
			"ssl_custom_options": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_default_certificate": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_hsts_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ssl_hsts_include_sub_domains": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_hsts_max_age": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(15768000),
			},
			"ssl_hsts_preload": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_max_version": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"ssl_min_version": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"stickiness_bytes_in_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"stickiness_bytes_out_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1m"),
			},
			"stickiness_conn_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_counter": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
			},
			"stickiness_counter_key": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("src"),
			},
			"stickiness_data_types": schema.SetAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue()),
			},
			"stickiness_expire": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("30m"),
			},
			"stickiness_http_err_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_http_req_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_length": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"stickiness_sess_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("10s"),
			},
			"stickiness_size": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("50k"),
			},
			"tuning_max_connections": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(-1),
			},
			"tuning_shards": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_client": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_http_keep_alive": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_http_req": schema.StringAttribute{
				MarkdownDescription: "TODO",
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
			"advertised_protocols": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "ba_advertised_protocols",
				Computed:            true,
			},
			"basic_auth_groups": schema.SetAttribute{
				MarkdownDescription: "basicAuthEnabled",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"basic_auth_users": schema.SetAttribute{
				MarkdownDescription: "basicAuthGroups",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"bind": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Required:            true,
			},
			"bind_options": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"connection_behaviour": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"custom_options": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"default_backend": schema.StringAttribute{
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
			"forward_for": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"http2enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"http2enabled_nontls": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"linked_actions": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"linked_cpu_affinity_rules": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"linked_errorfiles": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"logging_detailed_log": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"logging_dont_log_normal": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"logging_dont_log_null": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"logging_log_separate_errors": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"logging_socket_stats": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Required:            true,
			},
			"prometheus_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"prometheus_path": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_advanced_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_bind_options": schema.SetAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"ssl_certificates": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_cipher_list": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_cipher_suites": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_client_auth_cas": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_client_auth_crls": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_client_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_client_auth_verify": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_custom_options": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_default_certificate": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_hsts_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_hsts_include_sub_domains": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_hsts_max_age": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_hsts_preload": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_max_version": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ssl_min_version": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_bytes_in_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_bytes_out_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_conn_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_counter": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_counter_key": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_data_types": schema.SetAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"stickiness_expire": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_http_err_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_http_req_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_length": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_sess_rate_period": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_size": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_max_connections": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_shards": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_timeout_client": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_timeout_http_keep_alive": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_timeout_http_req": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
		},
	}
}

func convertHaproxyFrontendSchemaToStruct(d *HaproxyFrontendResourceModel) (*haproxy.Frontend, error) {
	return &haproxy.Frontend{
		AdvertisedProtocols:          api.SelectedMap(d.AdvertisedProtocols.ValueString()),
		BasicAuthEnabled:             tools.BoolToString(d.BasicAuthEnabled.ValueBool()),
		BasicAuthGroups:              tools.SetToString(d.BasicAuthGroups),
		BasicAuthUsers:               tools.SetToString(d.BasicAuthUsers),
		Bind:                         d.Bind.ValueString(),
		BindOptions:                  d.BindOptions.ValueString(),
		ConnectionBehaviour:          api.SelectedMap(d.ConnectionBehaviour.ValueString()),
		CustomOptions:                d.CustomOptions.ValueString(),
		DefaultBackend:               api.SelectedMap(d.DefaultBackend.ValueString()),
		Description:                  d.Description.ValueString(),
		Enabled:                      tools.BoolToString(d.Enabled.ValueBool()),
		ForwardFor:                   tools.BoolToString(d.ForwardFor.ValueBool()),
		Http2Enabled:                 tools.BoolToString(d.Http2Enabled.ValueBool()),
		Http2EnabledNontls:           tools.BoolToString(d.Http2EnabledNontls.ValueBool()),
		LinkedActions:                d.LinkedActions.ValueString(),
		LinkedCpuAffinityRules:       d.LinkedCpuAffinityRules.ValueString(),
		LinkedErrorfiles:             d.LinkedErrorfiles.ValueString(),
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
		SslCertificates:              api.SelectedMap(d.SslCertificates.ValueString()),
		SslCipherList:                d.SslCipherList.ValueString(),
		SslCipherSuites:              d.SslCipherSuites.ValueString(),
		SslClientAuthCAs:             api.SelectedMap(d.SslClientAuthCAs.ValueString()),
		SslClientAuthCRLs:            api.SelectedMap(d.SslClientAuthCRLs.ValueString()),
		SslClientAuthEnabled:         tools.BoolToString(d.SslClientAuthEnabled.ValueBool()),
		SslClientAuthVerify:          api.SelectedMap(d.SslClientAuthVerify.ValueString()),
		SslCustomOptions:             d.SslCustomOptions.ValueString(),
		SslDefaultCertificate:        api.SelectedMap(d.SslDefaultCertificate.ValueString()),
		SslEnabled:                   tools.BoolToString(d.SslEnabled.ValueBool()),
		SslHstsEnabled:               tools.BoolToString(d.SslHstsEnabled.ValueBool()),
		SslHstsIncludeSubDomains:     tools.BoolToString(d.SslHstsIncludeSubDomains.ValueBool()),
		SslHstsMaxAge:                tools.Int64ToStringNegative(d.SslHstsMaxAge.ValueInt64()),
		SslHstsPreload:               d.SslHstsPreload.ValueString(),
		SslMaxVersion:                api.SelectedMap(d.SslMaxVersion.ValueString()),
		SslMinVersion:                api.SelectedMap(d.SslMinVersion.ValueString()),
		StickinessBytesInRatePeriod:  d.StickinessBytesInRatePeriod.ValueString(),
		StickinessBytesOutRatePeriod: d.StickinessBytesOutRatePeriod.ValueString(),
		StickinessConnRatePeriod:     d.StickinessConnRatePeriod.ValueString(),
		StickinessCounter:            tools.Int64ToStringNegative(d.StickinessCounter.ValueInt64()),
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
		AdvertisedProtocols:          types.StringValue(d.AdvertisedProtocols.String()),
		BasicAuthEnabled:             types.BoolValue(tools.StringToBool(d.BasicAuthEnabled)),
		BasicAuthGroups:              tools.StringToSet(d.BasicAuthGroups),
		BasicAuthUsers:               tools.StringToSet(d.BasicAuthUsers),
		Bind:                         types.StringValue(d.Bind),
		BindOptions:                  types.StringValue(d.BindOptions),
		ConnectionBehaviour:          types.StringValue(d.ConnectionBehaviour.String()),
		CustomOptions:                types.StringValue(d.CustomOptions),
		DefaultBackend:               types.StringValue(d.DefaultBackend.String()),
		Description:                  types.StringValue(d.Description),
		Enabled:                      types.BoolValue(tools.StringToBool(d.Enabled)),
		ForwardFor:                   types.BoolValue(tools.StringToBool(d.ForwardFor)),
		Http2Enabled:                 types.BoolValue(tools.StringToBool(d.Http2Enabled)),
		Http2EnabledNontls:           types.BoolValue(tools.StringToBool(d.Http2EnabledNontls)),
		LinkedActions:                types.StringValue(d.LinkedActions),
		LinkedCpuAffinityRules:       types.StringValue(d.LinkedCpuAffinityRules),
		LinkedErrorfiles:             types.StringValue(d.LinkedErrorfiles),
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
		SslCertificates:              types.StringValue(d.SslCertificates.String()),
		SslCipherList:                types.StringValue(d.SslCipherList),
		SslCipherSuites:              types.StringValue(d.SslCipherSuites),
		SslClientAuthCAs:             types.StringValue(d.SslClientAuthCAs.String()),
		SslClientAuthCRLs:            types.StringValue(d.SslClientAuthCRLs.String()),
		SslClientAuthEnabled:         types.BoolValue(tools.StringToBool(d.SslClientAuthEnabled)),
		SslClientAuthVerify:          types.StringValue(d.SslClientAuthVerify.String()),
		SslCustomOptions:             types.StringValue(d.SslCustomOptions),
		SslDefaultCertificate:        types.StringValue(d.SslDefaultCertificate.String()),
		SslEnabled:                   types.BoolValue(tools.StringToBool(d.SslEnabled)),
		SslHstsEnabled:               types.BoolValue(tools.StringToBool(d.SslHstsEnabled)),
		SslHstsIncludeSubDomains:     types.BoolValue(tools.StringToBool(d.SslHstsIncludeSubDomains)),
		SslHstsMaxAge:                types.Int64Value(tools.StringToInt64(d.SslHstsMaxAge)),
		SslHstsPreload:               types.StringValue(d.SslHstsPreload),
		SslMaxVersion:                types.StringValue(d.SslMaxVersion.String()),
		SslMinVersion:                types.StringValue(d.SslMinVersion.String()),
		StickinessBytesInRatePeriod:  types.StringValue(d.StickinessBytesInRatePeriod),
		StickinessBytesOutRatePeriod: types.StringValue(d.StickinessBytesOutRatePeriod),
		StickinessConnRatePeriod:     types.StringValue(d.StickinessConnRatePeriod),
		StickinessCounter:            types.Int64Value(tools.StringToInt64(d.StickinessCounter)),
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
