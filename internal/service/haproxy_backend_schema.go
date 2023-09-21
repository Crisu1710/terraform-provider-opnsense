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

// HaproxyBackendResourceModel describes the resource data model.
type HaproxyBackendResourceModel struct {
	Algorithm                    types.String `tfsdk:"algorithm"`
	BaAdvertisedProtocols        types.String `tfsdk:"ba_advertised_protocols"`
	BasicAuthEnabled             types.Bool   `tfsdk:"basic_auth_enabled"`
	BasicAuthGroups              types.String `tfsdk:"basic_auth_groups"`
	BasicAuthUsers               types.String `tfsdk:"basic_auth_users"`
	CheckDownInterval            types.String `tfsdk:"check_down_interval"`
	CheckInterval                types.String `tfsdk:"check_interval"`
	CustomOptions                types.String `tfsdk:"custom_options"`
	Description                  types.String `tfsdk:"description"`
	Enabled                      types.Bool   `tfsdk:"enabled"`
	HealthCheck                  types.String `tfsdk:"health_check"`
	HealthCheckEnabled           types.Bool   `tfsdk:"health_check_enabled"`
	HealthCheckFall              types.String `tfsdk:"health_check_fall"`
	HealthCheckLogStatus         types.Bool   `tfsdk:"health_check_log_status"`
	HealthCheckRise              types.String `tfsdk:"health_check_rise"`
	Http2Enabled                 types.Bool   `tfsdk:"http2enabled"`
	Http2EnabledNontls           types.Bool   `tfsdk:"http2enabled_nontls"`
	LinkedActions                types.String `tfsdk:"linked_actions"`
	LinkedErrorfiles             types.String `tfsdk:"linked_errorfiles"`
	LinkedFcgi                   types.String `tfsdk:"linked_fcgi"`
	LinkedMailer                 types.String `tfsdk:"linked_mailer"`
	LinkedResolver               types.String `tfsdk:"linked_resolver"`
	LinkedServers                types.String `tfsdk:"linked_servers"`
	Mode                         types.String `tfsdk:"mode"`
	Name                         types.String `tfsdk:"name"`
	Persistence                  types.String `tfsdk:"persistence"`
	PersistenceCookiemode        types.String `tfsdk:"persistence_cookiemode"`
	PersistenceCookiename        types.String `tfsdk:"persistence_cookiename"`
	PersistenceStripquotes       types.Bool   `tfsdk:"persistence_stripquotes"`
	ProxyProtocol                types.String `tfsdk:"proxy_protocol"`
	RandomDraws                  types.Int64  `tfsdk:"random_draws"`
	ResolvePrefer                types.String `tfsdk:"resolve_prefer"`
	ResolverOpts                 types.String `tfsdk:"resolver_opts"`
	Source                       types.String `tfsdk:"source"`
	StickinessBytesInRatePeriod  types.String `tfsdk:"stickiness_bytes_in_rate_period"`
	StickinessBytesOutRatePeriod types.String `tfsdk:"stickiness_bytes_out_rate_period"`
	StickinessConnRatePeriod     types.String `tfsdk:"stickiness_conn_rate_period"`
	StickinessCookielength       types.String `tfsdk:"stickiness_cookielength"`
	StickinessCookiename         types.String `tfsdk:"stickiness_cookiename"`
	StickinessDataTypes          types.String `tfsdk:"stickiness_data_types"`
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
	TuningRetries                types.String `tfsdk:"tuning_retries"`
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
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("source"),
			},
			"ba_advertised_protocols": schema.StringAttribute{
				MarkdownDescription: "ba_advertised_protocols",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("h2,http11"),
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "basicAuthEnabled",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"basic_auth_groups": schema.StringAttribute{
				MarkdownDescription: "basicAuthGroups",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"basic_auth_users": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"check_down_interval": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"check_interval": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"custom_options": schema.StringAttribute{
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
			"health_check": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_check_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"health_check_fall": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"health_check_log_status": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"health_check_rise": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
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
			"linked_errorfiles": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_fcgi": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_mailer": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"linked_servers": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
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
			"persistence": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("sticktable"),
			},
			"persistence_cookiemode": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("piggyback"),
			},
			"persistence_cookiename": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("SRVCOOKIE"),
			},
			"persistence_stripquotes": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"proxy_protocol": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"random_draws": schema.Int64Attribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2),
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
			"source": schema.StringAttribute{
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
			"stickiness_cookielength": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"stickiness_cookiename": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"stickiness_data_types": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
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
			"stickiness_pattern": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("sourceipv4"),
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
			"tuning_caching": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tuning_defaultserver": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_httpreuse": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("safe"),
			},
			"tuning_noport": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tuning_retries": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_check": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_connect": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"tuning_timeout_server": schema.StringAttribute{
				MarkdownDescription: "TODO",
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
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"ba_advertised_protocols": schema.StringAttribute{
				MarkdownDescription: "ba_advertised_protocols",
				Computed:            true,
			},
			"basic_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: "basicAuthEnabled",
				Computed:            true,
			},
			"basic_auth_groups": schema.StringAttribute{
				MarkdownDescription: "basicAuthGroups",
				Computed:            true,
			},
			"basic_auth_users": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"check_down_interval": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"check_interval": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"custom_options": schema.StringAttribute{
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
			"health_check": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"health_check_enabled": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"health_check_fall": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"health_check_log_status": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"health_check_rise": schema.StringAttribute{
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
			"linked_errorfiles": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"linked_fcgi": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"linked_mailer": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"linked_resolver": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"linked_servers": schema.StringAttribute{
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
			"persistence": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"persistence_cookiemode": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"persistence_cookiename": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"persistence_stripquotes": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"proxy_protocol": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"random_draws": schema.Int64Attribute{
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
			"source": schema.StringAttribute{
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
			"stickiness_cookielength": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_cookiename": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"stickiness_data_types": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
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
			"tuning_caching": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_defaultserver": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_httpreuse": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_noport": schema.BoolAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_retries": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_timeout_check": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_timeout_connect": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
			"tuning_timeout_server": schema.StringAttribute{
				MarkdownDescription: "TODO",
				Computed:            true,
			},
		},
	}
}

func convertHaproxyBackendSchemaToStruct(d *HaproxyBackendResourceModel) (*haproxy.Backend, error) {
	return &haproxy.Backend{
		Algorithm:                    api.SelectedMap(d.Algorithm.ValueString()),
		BaAdvertisedProtocols:        api.SelectedMap(d.BaAdvertisedProtocols.ValueString()),
		BasicAuthEnabled:             tools.BoolToString(d.BasicAuthEnabled.ValueBool()),
		BasicAuthGroups:              d.BasicAuthGroups.ValueString(),
		BasicAuthUsers:               d.BasicAuthUsers.ValueString(),
		CheckDownInterval:            d.CheckDownInterval.ValueString(),
		CheckInterval:                d.CheckInterval.ValueString(),
		CustomOptions:                d.CustomOptions.ValueString(),
		Description:                  d.Description.ValueString(),
		Enabled:                      tools.BoolToString(d.Enabled.ValueBool()),
		HealthCheck:                  api.SelectedMap(d.HealthCheck.ValueString()),
		HealthCheckEnabled:           tools.BoolToString(d.HealthCheckEnabled.ValueBool()),
		HealthCheckFall:              d.HealthCheckFall.ValueString(),
		HealthCheckLogStatus:         tools.BoolToString(d.HealthCheckLogStatus.ValueBool()),
		HealthCheckRise:              d.HealthCheckRise.ValueString(),
		Http2Enabled:                 tools.BoolToString(d.Http2Enabled.ValueBool()),
		Http2EnabledNontls:           tools.BoolToString(d.Http2EnabledNontls.ValueBool()),
		LinkedActions:                d.LinkedActions.ValueString(),
		LinkedErrorfiles:             d.LinkedErrorfiles.ValueString(),
		LinkedFcgi:                   api.SelectedMap(d.LinkedFcgi.ValueString()),
		LinkedMailer:                 api.SelectedMap(d.LinkedMailer.ValueString()),
		LinkedResolver:               api.SelectedMap(d.LinkedResolver.ValueString()),
		LinkedServers:                d.LinkedServers.ValueString(),
		Mode:                         api.SelectedMap(d.Mode.ValueString()),
		Name:                         d.Name.ValueString(),
		Persistence:                  api.SelectedMap(d.Persistence.ValueString()),
		PersistenceCookiemode:        api.SelectedMap(d.PersistenceCookiemode.ValueString()),
		PersistenceCookiename:        d.PersistenceCookiename.ValueString(),
		PersistenceStripquotes:       tools.BoolToString(d.PersistenceStripquotes.ValueBool()),
		ProxyProtocol:                api.SelectedMap(d.ProxyProtocol.ValueString()),
		RandomDraws:                  tools.Int64ToString(d.RandomDraws.ValueInt64()),
		ResolvePrefer:                api.SelectedMap(d.ResolvePrefer.ValueString()),
		ResolverOpts:                 api.SelectedMap(d.ResolverOpts.ValueString()),
		Source:                       d.Source.ValueString(),
		StickinessBytesInRatePeriod:  d.StickinessBytesInRatePeriod.ValueString(),
		StickinessBytesOutRatePeriod: d.StickinessBytesOutRatePeriod.ValueString(),
		StickinessConnRatePeriod:     d.StickinessConnRatePeriod.ValueString(),
		StickinessCookielength:       d.StickinessCookielength.ValueString(),
		StickinessCookiename:         d.StickinessCookiename.ValueString(),
		StickinessDataTypes:          api.SelectedMap(d.StickinessDataTypes.ValueString()),
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
		TuningRetries:                d.TuningRetries.ValueString(),
		TuningTimeoutCheck:           d.TuningTimeoutCheck.ValueString(),
		TuningTimeoutConnect:         d.TuningTimeoutConnect.ValueString(),
		TuningTimeoutServer:          d.TuningTimeoutServer.ValueString(),
	}, nil
}

func convertHaproxyBackendStructToSchema(d *haproxy.Backend) (*HaproxyBackendResourceModel, error) {
	return &HaproxyBackendResourceModel{
		Algorithm:                    types.StringValue(d.Algorithm.String()),
		BaAdvertisedProtocols:        types.StringValue(d.BaAdvertisedProtocols.String()),
		BasicAuthEnabled:             types.BoolValue(tools.StringToBool(d.BasicAuthEnabled)),
		BasicAuthGroups:              types.StringValue(d.BasicAuthGroups),
		BasicAuthUsers:               types.StringValue(d.BasicAuthUsers),
		CheckDownInterval:            types.StringValue(d.CheckDownInterval),
		CheckInterval:                types.StringValue(d.CheckInterval),
		CustomOptions:                types.StringValue(d.CustomOptions),
		Description:                  types.StringValue(d.Description),
		Enabled:                      types.BoolValue(tools.StringToBool(d.Enabled)),
		HealthCheck:                  types.StringValue(d.HealthCheck.String()),
		HealthCheckEnabled:           types.BoolValue(tools.StringToBool(d.HealthCheckEnabled)),
		HealthCheckFall:              types.StringValue(d.HealthCheckFall),
		HealthCheckLogStatus:         types.BoolValue(tools.StringToBool(d.HealthCheckLogStatus)),
		HealthCheckRise:              types.StringValue(d.HealthCheckRise),
		Http2Enabled:                 types.BoolValue(tools.StringToBool(d.Http2Enabled)),
		Http2EnabledNontls:           types.BoolValue(tools.StringToBool(d.Http2EnabledNontls)),
		LinkedActions:                types.StringValue(d.LinkedActions),
		LinkedErrorfiles:             types.StringValue(d.LinkedErrorfiles),
		LinkedFcgi:                   types.StringValue(d.LinkedFcgi.String()),
		LinkedMailer:                 types.StringValue(d.LinkedMailer.String()),
		LinkedResolver:               types.StringValue(d.LinkedResolver.String()),
		LinkedServers:                types.StringValue(d.LinkedServers),
		Mode:                         types.StringValue(d.Mode.String()),
		Name:                         types.StringValue(d.Name),
		Persistence:                  types.StringValue(d.Persistence.String()),
		PersistenceCookiemode:        types.StringValue(d.PersistenceCookiemode.String()),
		PersistenceCookiename:        types.StringValue(d.PersistenceCookiename),
		PersistenceStripquotes:       types.BoolValue(tools.StringToBool(d.PersistenceStripquotes)),
		ProxyProtocol:                types.StringValue(d.ProxyProtocol.String()),
		RandomDraws:                  types.Int64Value(tools.StringToInt64(d.RandomDraws)),
		ResolvePrefer:                types.StringValue(d.ResolverOpts.String()),
		ResolverOpts:                 types.StringValue(d.ResolverOpts.String()),
		Source:                       types.StringValue(d.Source),
		StickinessBytesInRatePeriod:  types.StringValue(d.StickinessBytesInRatePeriod),
		StickinessBytesOutRatePeriod: types.StringValue(d.StickinessBytesOutRatePeriod),
		StickinessConnRatePeriod:     types.StringValue(d.StickinessConnRatePeriod),
		StickinessCookielength:       types.StringValue(d.StickinessCookielength),
		StickinessCookiename:         types.StringValue(d.StickinessCookiename),
		StickinessDataTypes:          types.StringValue(d.StickinessDataTypes.String()),
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
		TuningRetries:                types.StringValue(d.TuningRetries),
		TuningTimeoutCheck:           types.StringValue(d.TuningTimeoutCheck),
		TuningTimeoutConnect:         types.StringValue(d.TuningTimeoutConnect),
		TuningTimeoutServer:          types.StringValue(d.TuningTimeoutServer),
	}, nil
}
