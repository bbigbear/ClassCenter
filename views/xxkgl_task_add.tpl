<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>新建任务</title>
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
	<label class="layui-form-label">任务编号</label>
	<div class="layui-inline">			    	
		<div class="layui-input-inline" style="width: 100px;">
      		<input type="text" name="TaskId" id="TaskId" autocomplete="off" class="layui-input">
    	</div>
	</div>
  </div>
  <div class="layui-form-item">
	<label class="layui-form-label">任务名称</label>
	<div class="layui-input-inline" >
		<input type="text" name="TaskName" id="TaskName" autocomplete="off" class="layui-input">	    	
	</div>
	<div class="layui-form-mid layui-word-aux">*只读</div>
  </div>  
  <div class="layui-form-item">
	    <label class="layui-form-label">所属学期</label>
	    <div class="layui-input-inline" >
		    <select name="Term" id="Term" lay-filter="style_select">
			    <option value="未选择" > 未选择</option>
				<option value="上学期" > 上学期</option>
				<option value="下学期" > 下学期</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	<label class="layui-form-label">开始学期</label>
	<div class="layui-input-inline" >
		<input type="text" name="date" id="StartTime" autocomplete="off" class="layui-input">	    	
	</div>
  </div> 
  <div class="layui-form-item">
	<label class="layui-form-label">结束学期</label>
	<div class="layui-input-inline" >
		<input type="text" name="date" id="EndTime" autocomplete="off" class="layui-input">	    	
	</div>
  </div> 
  <div class="layui-form-item">
	    <label class="layui-form-label">对教师是否启动</label>
	    <div class="layui-input-inline" >
		    <select name="TeacherEnable" id="TeacherEnable" lay-filter="style_select">
			    <option value="未选择" > 未选择</option>
				<option value="启动" > 启动</option>
				<option value="不启动" > 不启动</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label">对学生是否启动</label>
	    <div class="layui-input-inline" >
		    <select name="StudentEnable" id="StudentEnable" lay-filter="style_select">
			    <option value="未选择" > 未选择</option>
				<option value="启动" > 启动</option>
				<option value="不启动" > 不启动</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
    <div class="layui-input-block">
      <button class="layui-btn layui-btn-primary" id="add">添加</button>
    </div>
  </div>
</form>

<br><br><br>

<script src="/static/layui.js"></script>
<!-- <script src="../build/lay/dest/layui.all.js"></script> -->
<script>
layui.use(['form','laydate','upload','jquery','layedit','element'], function(){
  var form = layui.form
  ,laydate=layui.laydate
  ,upload = layui.upload
  , $ = layui.jquery
  ,layedit=layui.layedit
  ,element=layui.element;


	laydate.render({
	    elem: '#StartTime'
	    ,type: 'date'
	  });
	
	laydate.render({
	    elem: '#EndTime'
	    ,type: 'date'
	  });

	//数据上传
	function uploadData(){
		var TeacherEnable=$("#TeacherEnable").val()
		var StudentEnable=$("#StudentEnable").val()
		var Term=$("#Term").val()
		var data={
				'TaskId':$("#TaskId").val(),
				'Name':$("#TaskName").val(),
				'Term':Term,
				'StartTime':new Date($("#StartTime").val()),
				'EndTime':new Date($("#EndTime").val()),
				'TeacherEnable':TeacherEnable,
				'StudentEnable':StudentEnable
			};
			console.log(data)
			if(TeacherEnable=="请选择"||StudentEnable=="请选择"||Term=="请选择"){
				alert("下拉列表不能为空")
			}else{
				//发布
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/xxkgl/task/add_action",
					data:JSON.stringify(data),
					async:false,
					error:function(request){
						alert("post error")						
					},
					success:function(res){
						if(res.code==200){
							alert("保存成功")
							window.location.reload();		
						}else{
							alert("保存失败")
						}						
					}
			  	});	
			}
	}
	
	$('#add').on('click',function(){
		uploadData()
		return false;
	});
	
});
</script>

</body>
</html>