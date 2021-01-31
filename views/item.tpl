<!DOCTYPE html>
<html>
  <head>
    <title>
    {{if .item}}
      修改事项
    {{else}}
      新增事项
    {{end}}
    </title>
    <link rel="stylesheet" type="text/css" href="/node_modules/bootstrap/dist/css/bootstrap.css" />
    <link rel="stylesheet" type="text/css" href="/static/css/common.css" />
  </head>
  <body>
    {{if .item}}
    <div class="box-form">
      <form>
        <div class="form-group">
          <label>标题：</label>
          <input type="text" class="form-control" name="title" value="{{.item.Title}}" />
        </div>
        <div class="form-group">
          <label>详情：</label>
          <textarea name="detail" class="form-control" rows="10">{{.item.Detail}}</textarea>
        </div>
        <div class="form-group">
          <label>期限：</label>
          <input type="datetime-local" class="form-control" name="deadline" step="01" value="{{.item.Deadline}}"/>
        </div>
        <div class="form-group">
          <label>参与人：</label>
          <input type="text" class="form-control" name="people" value="{{call $.JoinPeople .item.People}}"/>
        </div>
        <div class="text-right">
          <button type="submit" class="btn btn-primary" formmethod="post">提交</button>
          <button type="reset" class="btn btn-danger btn-cancel">取消</button>
        </div>
      </form>
    </div>
    {{else}}
    <div class="box-form">
      <form>
        <div class="form-group">
          <label>标题：</label>
          <input type="text" class="form-control" name="title" />
        </div>
        <div class="form-group">
          <label>详情：</label>
          <textarea name="detail" class="form-control" rows="10"></textarea>
        </div>
        <div class="form-group">
          <label>期限：</label>
          <input type="datetime-local" class="form-control" name="deadline" step="01"/>
        </div>
        <div class="form-group">
          <label>参与人：</label>
          <input type="text" class="form-control" name="people" />
        </div>
        <div class="text-right">
          <button type="submit" class="btn btn-primary" formmethod="post">提交</button>
          <button type="reset" class="btn btn-danger btn-cancel">取消</button>
        </div>
      </form>
    </div>
    {{end}}
    <script src="/node_modules/jquery/dist/jquery.js"></script>
    <script src="/node_modules/popper.js/dist/popper.js" type="module"></script>
    <script src="/node_modules/bootstrap/dist/js/bootstrap.js"></script>
    <script type="text/javascript">
      $(function() {
        $(".btn-cancel").click(function() {
          $(window).attr("location", "/");
        });
      });
    </script>
  </body>
</html>