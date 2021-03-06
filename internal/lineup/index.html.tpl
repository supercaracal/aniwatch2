<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
  <link rel="shortcut icon" type="image/x-icon" href="/{{.Data.AppName}}/favicon.ico" />
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css" integrity="sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2" crossorigin="anonymous">
  <link rel="stylesheet" media="all" href="/{{.Data.AppName}}/style.css" />
  <title>{{.Title}}</title>
</head>
<body>
  <div class="content-wrap">
    <div class="navbar navbar-dark bg-dark global-header">
      <a class="navbar-brand" href="javascript:;">{{.Title}}</a>
    </div>
    <div class="container-fluid">
      <div class="row mt-3">
        <div class="col-md-12">
          <table class="table table-bordered">
            <thead>
              <tr>
                <th><span>&#10004;{{.LineupCount}}</span></th>
                {{range .Data.SlotOfDay}}
                  <th scope="col"><span class="badge badge-{{.Color}}">{{.Name}}</span></th>
                {{end}}
              </tr>
            </thead>
            <tbody>
              {{range .Data.DayOfWeek}}
                <tr>
                  <th scope="row"><span class="badge badge-pill badge-light">{{.Name}}</span></th>
                  {{with $slots := index $.LineupsPerDaySlot .ID}}
                    {{range $.Data.SlotOfDay}}
                      <td>
                        {{with $lineups := index $slots .ID}}
                          <ul class="list-unstyled">
                            {{range $lineups}}
                              <li>
                                <strong>{{.Title}}</strong>
                                <span class="text-muted">
                                  {{.Start}}&nbsp;{{.Channel}}
                                </span>
                              </li>
                            {{end}}
                          </ul>
                        {{else}}
                          <span class="text-muted">--</span>
                        {{end}}
                      </td>
                    {{end}}
                  {{end}}
                </tr>
              {{end}}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  <div class="global-footer">
    <p class="text-md-center text-muted">
      <span>{{.Quarter}}</span>
      <a class="ml-1" href="https://github.com/{{.Data.AuthorID}}/{{.Data.AppName}}">{{.Data.AuthorName}}</a>
    </p>
  </div>
</body>
