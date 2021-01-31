<!DOCTYPE html>
<html>
<head>
  <title>历史记录</title>
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
    <div class="box">
      <div class="content" id="item-{{.Id}}">
        <h2>{{.Title}}</h2>
      </div>
    </div>
    {{end}}
  </div>
  <script src="/node_modules/jquery/dist/jquery.js"></script>
  <script src="/node_modules/popper.js/dist/popper.js" type="module"></script>
  <script src="/node_modules/bootstrap/dist/js/bootstrap.js"></script>
  <script src="/static/js/reload.min.js"></script>
  <script type="text/javascript">
    $(function() {
      console.log({{.items}});
      var items = {{.items}};
      var history = {};
      for (var i = 0; i < items.length; i++) {
        history[items[i].Id] = items[i];
      }
      $(".box").mouseover(function() {
        var id = $($(this).children()[0]).attr("id").split("-")[1];
        $("#title").text("标题：" + history[id].Title);
        $("#detail").text("详情：" + history[id].Detail);
        $("#deadline").text("期限：" + history[id].Deadline);
        $("#people").text("参与人：" + history[id].People.join(" "));
      });
    });
  </script>
</body>
</html>