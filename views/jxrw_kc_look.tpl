<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>课程查看</title>
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <meta name="apple-mobile-web-app-status-bar-style" content="black"> 
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="format-detection" content="telephone=no">
  <link rel="stylesheet" href="/static/css/layui.css">

<style>
body{padding: 10px;}
</style>
</head>
<body>
<form class="layui-form layui-form-pane1" action="">
  <div class="layui-inline">
    <label class="layui-form-label">课程名称</label>
    <div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-inline">
    <label class="layui-form-label">学期</label>
    <div class="layui-input-inline" style="width: 150px;">
        <select name="Major" id="Major" lay-filter="style_select">
			    <option value="请选择" > 请选择</option>
				<option value="上学期" > 上学期</option>
				<option value="下学期" > 下学期</option>
	    </select>
    </div>
  </div>
  <div class="layui-form-item" style="padding-top:10px;padding-left:50px;">
    <div class="layui-inline">
    	<button class="layui-btn" id="query">查询</button>
  	</div>
  </div>
</form>
	<table class="layui-table" id="list" lay-filter="announcement" lay-size="sm"></table>
<script src="/static/layui.js"></script>
<script src="http://code.jquery.com/jquery-2.1.1.min.js"></script>
<!-- <script src="../build/lay/dest/layui.all.js"></script> -->

<script type="text/javascript">
layui.use(['form','laydate','upload','jquery','layedit','element','table','laytpl'], function(){
  var form = layui.form
  ,laydate=layui.laydate
  ,upload = layui.upload
  , $ = layui.jquery
  ,layedit=layui.layedit
  ,element=layui.element
  ,table=layui.table
  ,laytpl = layui.laytpl;
	 //table 渲染
	  table.render({
	    elem: '#list'
	    ,height: 315
	    //,url: '/v1/jxkc/getdata'//数据接口
		,data:[<.m>]
	    //,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[
	      {field:'Id', title:'序号', width:80}
		  ,{field:'Term',  title:'执行学期', width:120}
	      ,{field:'CourseId',  title:'课程号', width:80}
		  ,{field:'CourseName',  title:'课程名称', width:80}
		  ,{field:'Core',  title:'是否核心课程', width:120}
	    ]]
	  });	
	$('#query').on('click',function(){
		//alert("点击查询")
		var CourseName=$("#CourseName").val()
		table.reload('listReload',{
			where:{
				CourseName:CourseName,
			}
		})		
		return false;
	})
	$('#add').on('click',function(){
		var checkStatus=table.checkStatus('listReload')
		var data=checkStatus.data
		console.log(data);
		if(data.length==0){
			alert("请选择要添加的课程")
		}else if(data.length>1){
			alert("只能选择一门课程")
		}else{
			parent.$('#CourseId').val(data[0].CourseId);
			parent.$('#CourseName').val(data[0].CourseName)
			var index = parent.layer.getFrameIndex(window.name);  
			parent.layer.close(index);	
		}
		return false;
	})
	
   

	
});
</script>

</body>
</html>