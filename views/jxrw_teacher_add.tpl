<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>选择课程</title>
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
    <label class="layui-form-label">教师姓名</label>
    <div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item" style="padding-top:10px;padding-left:50px;">
    <div class="layui-inline">
    	<button class="layui-btn" id="query">查询</button>
		<button class="layui-btn" id="add">确认选择</button>
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
	//自动加载
	$(function(){
				
	});	
	var data=[{"TeacherId":"教师001","TeacherName":"张三"},{"TeacherId":"教师002","TeacherName":"李四"}]
	 //table 渲染
	  table.render({
	    elem: '#list'
	    ,height: 315
	    //,url: '/v1/jxkc/getdata'//数据接口
		,data:data
	    //,page: true //开启分页
		,id: 'listReload'
	    ,cols: [[
		  {type:'checkbox', fixed: 'left'}	
	      ,{field:'TeacherId', title:'教师工号', width:120}
		  ,{field:'TeacherName',  title:'教师姓名', width:120}
	    ]]
	  });	
	$('#query').on('click',function(){
		//alert("点击查询")
		var TeacherName=$("#TeacherName").val()
		table.reload('listReload',{
			where:{
				TeacherName:TeacherName,
			}
		})		
		return false;
	})
	$('#add').on('click',function(){
		var checkStatus=table.checkStatus('listReload')
		var data=checkStatus.data
		console.log(data);
		if(data.length==0){
			alert("请选择要添加的教师")
		}else if(data.length>1){
			alert("只能选择一个教师")
		}else{
			parent.$('#TeacherId').val(data[0].TeacherId);
			parent.$('#TeacherName').val(data[0].TeacherName)
			var index = parent.layer.getFrameIndex(window.name);  
			parent.layer.close(index);	
		}
		return false;
	})
	
   

	
});
</script>

</body>
</html>