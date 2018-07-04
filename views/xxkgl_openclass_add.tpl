<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>新增申请</title>
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
<form class="layui-form" action="">
  <div class="layui-form-item">
	<label class="layui-form-label">课程号</label>
	<div class="layui-inline">			    	
		<div class="layui-input-inline" style="width: 100px;">
      		<input type="text" name="CourseId" id="CourseId" autocomplete="off" class="layui-input">
    	</div>
	</div>
  </div>
  <div class="layui-form-item">
	<label class="layui-form-label">课程名称</label>
	<div class="layui-inline">			    	
		<div class="layui-input-inline" style="width: 100px;">
      		<input type="text" name="CourseName" id="CourseName" autocomplete="off" class="layui-input">
    	</div>
	</div>
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label">开课年级</label>
	    <div class="layui-input-inline" >
		    <select name="Year" id="Year" lay-filter="style_select">
			    <option value="未选择" > 未选择</option>
				<option value="2018" > 2018</option>
				<option value="2017" > 2017</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	<label class="layui-form-label">周课时</label>
	<div class="layui-input-inline" >
		<input type="text" name="WeekTime" id="WeekTime" autocomplete="off" class="layui-input">	    	
	</div>
	<div class="layui-form-mid layui-word-aux">*只读</div>
  </div>  
  <div class="layui-form-item">
	    <label class="layui-form-label">学生性别</label>
	    <div class="layui-input-inline" >
		    <select name="Sex" id="Sex" lay-filter="style_select">
			    <option value="未选择" > 未选择</option>
				<option value="男" > 男</option>
				<option value="女" > 女</option>
				<option value="男/女" > 男/女</option>
	      	</select>
		</div>		
  </div>
 <div class="layui-form-item">
	<label class="layui-form-label">开始日期</label>
	<div class="layui-input-inline" >
		<input type="text" name="date" id="StartTime" autocomplete="off" class="layui-input">	    	
	</div>
  </div> 
 <div class="layui-form-item">
	<label class="layui-form-label">截止日期</label>
	<div class="layui-input-inline" >
		<input type="text" name="date" id="EndTime" autocomplete="off" class="layui-input">	    	
	</div>
  </div> 
 <div class="layui-form-item">
	<label class="layui-form-label">申请人</label>
	<div class="layui-inline">			    	
		<div class="layui-input-inline" style="width: 100px;">
      		<input type="text" name="TeacherName" id="TeacherName" autocomplete="off" class="layui-input">
    	</div>
		<div class="layui-input-inline" style="width: 100px;">
      		<button class="layui-btn layui-btn-primary" id="choose">选择教师</button>
    	</div>
	</div>
  </div>
  <div class="layui-form-item">
	<label class="layui-form-label">备注</label>
	<div class="layui-input-block">
      <textarea class="layui-textarea" name="Remark" id="Remark"></textarea>
    </div>
  </div> 
  <div class="layui-form-item">
    <div class="layui-inline layui-layout-right">
      <button class="layui-btn layui-btn-primary" id="add">保存</button>
    </div>
  </div>
</form>

<br><br><br>

<script src="http://code.jquery.com/jquery-2.1.1.min.js"></script>
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
	
	//文本域
	var index=layedit.build('detail',{
		hideTool:['image','face']
	});

	//数据上传
	function uploadData(){
		var Year=$("#Year").val()
		var Sex=$("#Sex").val()
		var data={
			'CourseId':$("#CourseId").val(),
			'CourseName':$("#CourseName").val(),
			'Year':Year,
			'WeekTime':parseInt($("#WeekTime").val()),
			'Sex':Sex,
			'StartTime':new Date($("#StartTime").val()),
			'EndTime':new Date($("#EndTime").val()),
			'Applicant':$("#TeacherName").val(),
			'Remark':$("#Remark").val()
			};
			console.log(data)
			if(Year=="请选择"||Sex=="请选择"){
				alert("不能为空")
			}else{
				//发布
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/xxkgl/openclass_add_action",
					data:JSON.stringify(data),
					async:false,
					error:function(request){
						alert("post error")						
					},
					success:function(res){
						if(res.code==200){
							alert("保存成功")	
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
	
	$('#choose').on('click',function(){
		layer.open({
			type: 2,
			title: '课程选择',
			//closeBtn: 0, //不显示关闭按钮
			shadeClose: true,
			shade: false,
			area: ['700px', '500px'],
			anim: 2,
			content: ['/v1/jxjh/jxrw/teacher_add'], //iframe的url，no代表不显示滚动条
			cancel: function(index, layero){			  
				layer.close(index)
				window.location.reload();
				return false; 
			},
		});				
		return false;
	});

	
});
</script>

</body>
</html>