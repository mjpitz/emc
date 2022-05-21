<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Service Catalog</title>

  <style>
    body {
      font-family: sans-serif;
      background-color: #eaeaea;
      padding: 60px 0;
      margin: 0;
      line-height: 1.5rem;
    }

    div.catalog {
      max-width: 1600px;
      min-width: 1200px;
      margin: 0 auto;
    }

    div.row {
      display: flex;
    }

    div.col {
      flex: auto;
    }

    div.col-50 {
      max-width: 50%;
      min-width: 50%;
    }

    div.col-65 {
      max-width: 65%;
      min-width: 65%;
    }

    div.col-35 {
      max-width: 35%;
      min-width: 35%;
    }

    div.catalog h1 {
      text-align: center;
      font-size: 3rem;
      margin-bottom: 35px;
    }

    div.catalog div.service {
      background-color: #fff;
      border: 1px solid #aaa;
      padding: 20px 26px;
      border-radius: 5px;
      margin: 9px 18px;
    }

    div.catalog div.service img.logo {
      max-height: 48px;
      margin-right: 10px;
    }

    div.catalog div.service h2.name {
      display: inline-block;
      vertical-align: top;
      margin: 13px 0;
    }

    div.catalog div.service p.url {}
    div.catalog div.service p.description {}

    div.catalog div.service div.metadata {
      margin: 10px 0;
    }

    div.catalog div.service div.link-group {}
    div.catalog div.service div.link-group h3.label {}
    div.catalog div.service div.link-group li.link {}
  </style>
</head>

<body>
  <div class="catalog">
    {{- range $i, $service := .Services }}
    {{- if $i | mod 2 | eq 0 }}<div class="row">{{- end }}
    <div class="col col-50">
      <div class="service">
        <div class="row">
          <div class="col col-65" style="padding-right: 10px">
            <div>
              {{- if $service.LogoURL }}
              <img class="logo" src="{{ $service.LogoURL }}" />
              {{- end }}

              <h2 class="name">{{ $service.Label }}</h2>
            </div>

            {{- if $service.URL }}
            <p class="url"><a href="{{ $service.URL }}">{{ $service.URL }}</a></p>
            {{- end }}

            {{- if $service.Description }}
            <p class="description">{{ $service.Description }}</p>
            {{- end }}

            {{- range $kv := $service.Metadata }}
            <div class="metadata row">
              <div class="col col-35"><b>{{ $kv.Key }}</b></div>
              <div class="col col-65">{{ $kv.Value }}</div>
            </div>
            {{- end }}
          </div>

          <div class="col col-35">
            {{- range $group := $service.LinkGroups }}
            <div class="link-group">
              <h3 class="label">{{ $group.Label }}</h3>
              <ul>
                {{- range $link := $group.Links }}
                <li class="link">
                  <a href="{{ $link.URL }}">{{ $link.Label }}</a>
                </li>
                {{- end }}
              </ul>
            </div>
            {{- end }}
          </div>
        </div>
      </div>
    </div>
    {{- if $i | mod 2 | eq 1 }}</div>{{- end }}
    {{- end }}
  </div>
</body>
</html>
