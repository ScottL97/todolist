<!DOCTYPE html>
<html>
<head>
  <title>todolist</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <!--<link rel="stylesheet" type="text/css" href="/static/css/common.css" />-->
  <!--<link rel="stylesheet" type="text/css" href="/node_modules/bootstrap/dist/css/bootstrap.css" />-->
  <script type="text/javascript">
    document.write('<link rel="stylesheet" href="/node_modules/bootstrap/dist/css/bootstrap.css?time=' +
            new Date().getTime() + '"/>')
    document.write('<link rel="stylesheet" href="/static/css/common.css?time=' + new Date().getTime() + '"/>')
  </script>
</head>
<body>
  {{template "header.tpl" .}} <!--.用于将上下文传给子模板，让子模板的.和该模板文件一致-->
  <div class="container">
    {{range .items}}
    <div class="box{{call $.CheckDeadline .Deadline}}">
      <a href="/item/{{.Id}}">
        <div class="content">
          <h2>{{.Title}}</h2>
          <p>{{call $.GetRemainTime .Deadline}}</p>
        </div>
      </a>
      <div class="button-group">
        <button class="btn-light">详情</button>
        <button class="btn-success">完成</button>
      </div>
    </div>
    {{end}}
    <div class="box">
      <a href="/item">
        <div class="content">
          <!--<img src="/static/img/add.png" alt="添加事项" />-->
          <span class="add">+</span>
        </div>
      </a>
    </div>
  </div>
  <script src="/node_modules/jquery/dist/jquery.js"></script>
  <script src="/node_modules/popper.js/dist/popper.js" type="module"></script>
  <script src="/node_modules/bootstrap/dist/js/bootstrap.js"></script>
  <!--<script src="/static/js/reload.min.js"></script>-->
  <script type="text/javascript">
    $(function(){
      $("button").click(function(){
        var href = $(this).parent().prev().attr("href");
        switch ($(this).attr("class")) {
          case "btn-light": { // 详情
            $(window).attr("location", href);
          } break;
          case "btn-success": { // 删除，加入历史记录
            var id = href.split("/")[2];
            var item = $(this)
            $.post("/history", {"id": id}, function(data){
              // 更新左侧的待办、拖延、已完成
              $(window).attr("location", "/");
            });
          } break;
          default: break;
        }
      });
    });
  </script>
</body>
</html>