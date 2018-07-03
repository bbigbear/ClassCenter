<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>选课任务管理</title>
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
  <div class="layui-form-item">
    <div class="layui-inline" style="padding-left:50px;">
    	<button class="layui-btn" id="add">增加任务</button>
		<button class="layui-btn" id="delete">删除</button>
  	</div>
  </div>
</form>
<br><br>

	<table id="list" lay-filter="announcement" style="width:auto;"></table>
	<script type="text/html" id="barDemo">
		<a class="layui-btn layui-btn-normal layui-btn-xs" lay-event="edit">修改</a>
	</script>
<script src="/static/layui.js"></script>
<!-- <script src="../build/lay/dest/layui.all.js"></script> -->

<script>
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
	    ,url: '/v1/xxkgl/task/getdata'//数据接口
	    //,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[   
	      {field:'TaskId', title:'任务编号', width:120}
		  ,{field:'Name',  title:'任务名称', width:120}
	      ,{field:'Term',  title:'所属学期', width:120}
		  ,{field:'StartTime',  title:'任务开始日期', width:120}
		  ,{field:'EndTime',  title:'任务结束日期', width:120}
		  ,{field:'TeacherEnable',  title:'是否对教师启动', width:120}
		  ,{field:'StudentEnable',  title:'是否对学生启动', width:120}
		  ,{fixed: 'right', title:'操作',width:80, align:'center', toolbar: '#barDemo'}
	    ]]
	  });
	//监听工具条
		table.on('tool(announcement)', function(obj){ //注：tool是工具条事件名，test是table原始容器的属性 lay-filter="对应的值"
		    var data = obj.data //获得当前行数据
		    ,layEvent = obj.event; //获得 lay-event 对应的值
		    if(layEvent === 'edit'){
		      //layer.msg('查看操作');		
			  layer.open({
			  type: 2,
			  title: '修改任务',
			  //closeBtn: 0, //不显示关闭按钮
			  shadeClose: true,
			  shade: false,
			  area: ['893px', '600px'],
			 // offset: 'rb', //右下角弹出
			  //time: 2000, //2秒后自动关闭
			  maxmin: true,
			  anim: 2,
			  content: ['/v1/xxkgl/task/edit?id='+data.Id], //iframe的url，no代表不显示滚动条
			  cancel: function(index, layero){			  
				layer.close(index)
				window.location.reload();
			  	return false; 
			  },
		});
	    }
	  });
	$('#add').on('click',function(){
		layer.open({
			type: 2,
			title: '新建任务',
			//closeBtn: 0, //不显示关闭按钮
			shadeClose: true,
			shade: false,
			area: ['700px', '500px'],
			anim: 2,
			content: ['/v1/xxkgl/task/add'], //iframe的url，no代表不显示滚动条
			cancel: function(index, layero){			  
				layer.close(index)
				window.location.reload();
				return false; 
			},
		});	
		return false
	})
	
});
</script>

</body>
</html>