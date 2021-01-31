<div class="header">
  <a href="/"><h1>Welcome to todolist</h1></a>
  {{if .now}}
  <div class="summary">
    <p class="todo-num">
      待办：<span id="todo-num">{{.count}}</span>
    </p>
    <p class="delay-time">
      拖延：<span id="delay-time">{{.delayDay}} 天 {{.delayHour}} 时 {{.delayMinute}} 分</span>
    </p>
    <p class="history-num">
      已完成：<a href="/history"><span id="history-num">{{.historyCount}}</span></a>
    </p>
    <p>
      最近更新：{{.now}}
    </p>
    <p>
      Website: <a href="http://{{.website}}">{{.website}}</a>
    </p>
    <p>
      Contact me: <a class="email" href="mailto:{{.eail}}">{{.email}}</a>
    </p>
  </div>
  {{else}}
  <div class="summary">
    <p id="title" class="history-info"></p>
    <p id="detail" class="history-info"></p>
    <p id="deadline" class="history-info"></p>
    <p id="people" class="history-info"></p>
  </div>
  {{end}}
</div>