# update-time: {{.Domain.UpdatedAt}}
# domain-name : {{.Domain.DomainName}}
{{range $key, $value := .Pools}}
upstream {{ $.Domain.DomainName }}_{{ $key }} { {{ range $value.PoolNodes}}
    {{ if eq $value.Pool.Policy "ip_hash" }}ip_hash;{{end}}
    server {{.IP}}:{{.Port}} max_fails={{.MaxFailed}} weight={{.Weight}} fail_timeout={{.TimeOut}};{{ end }}
    keepalive {{ $value.Pool.Keepalive }};
}
{{end}}

{{ if  .Domain.UseHttps  }}
server {
    listen {{.Domain.HttpsPort}} ssl;
    server_name {{.Domain.Domain}};
    ssl_ciphers "EECDH+AESGCM:EDH+AESGCM:ECDHE-RSA-AES128-GCM-SHA256:AES256+EECDH:DHE-RSA-AES128-GCM-SHA256:AES256+EDH:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA256:ECDHE-RSA-AES256-SHA:ECDHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA256:DHE-RSA-AES128-SHA256:DHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA:ECDHE-RSA-DES-CBC3-SHA:EDH-RSA-DES-CBC3-SHA:AES256-GCM-SHA384:AES128-GCM-SHA256:AES256-SHA256:AES128-SHA256:AES256-SHA:AES128-SHA:DES-CBC3-SHA:HIGH:!aNULL:!eNULL:!EXPORT:!DES:!MD5:!PSK:!RC4";
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2 SSLv3;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL2:100m;
    ssl_session_timeout 10m;
    ssl_stapling on;

    ssl_certificate /etc/nginx/conf.d/cert/{{.Domain.HttpsCert}}.pem;
    ssl_certificate_key /etc/nginx/conf.d/cert/{{.Domain.HttpsCert}}.key;

    proxy_set_header Host $host;
    location /{
        set $ups {{.Domain.DomainName}}_{{.Domain.PoolCluster.PoolName}};
        {{ range .Rules}}
        location ^~ {{.Rule.Path}} {
            set $ups {{$.Domain.DomainName}}_{{.Pool.PoolName}};
            proxy_pass http://$ups;
            {{ range .Ins }}{{ if eq .InsType  "proxy_pass" }}{{ else }}{{ .InsType }}   {{ .InsValue }};{{ end }}{{ end }}
        }
        {{ end }}
        error_page                    500 502 503 504  /50x.html;
        location = /50x.html {
            root                        html;
        }

        if ( !-f $request_filename ) {
            proxy_pass                http://$ups;
            break;
        }
    }
}
{{end}}


server {
    listen {{.Domain.Port}};
    server_name {{.Domain.Domain}};
    proxy_set_header Host $host;
    location /{
        set $ups {{.Domain.DomainName}}_{{.Domain.PoolCluster.PoolName}};
        {{ range .Rules}}
        location ^~ {{.Rule.Path}} {
            set $ups {{$.Domain.DomainName}}_{{.Pool.PoolName}};
            proxy_pass http://$ups;
            {{ range .Ins }}{{ if eq .InsType  "proxy_pass" }}{{ else }}{{ .InsType }}   {{ .InsValue }};{{ end }}{{ end }}
        }
        {{ end }}
        error_page                    500 502 503 504  /50x.html;
        location = /50x.html {
            root                        html;
        }

        if ( !-f $request_filename ) {
            proxy_pass                http://$ups;
            break;
        }
    }
}

