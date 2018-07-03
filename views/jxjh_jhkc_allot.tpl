<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>分配计划课程</title>
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
		<div class="layui-input-inline" style="width: 100px;">
      		<button class="layui-btn layui-btn-primary" id="choose">选择课程</button>
    	</div>
	</div>
  </div>
  <div class="layui-form-item">
	<label class="layui-form-label">课程名称</label>
	<div class="layui-input-inline" >
		<input type="text" name="CourseName" id="CourseName" autocomplete="off" class="layui-input">	    	
	</div>
	<div class="layui-form-mid layui-word-aux">*修改时候只读</div>
  </div>  
  <div class="layui-form-item">
	    <label class="layui-form-label">执行学期</label>
	    <div class="layui-input-inline" >
		    <select name="Term" id="Term" lay-filter="style_select">
			    <option value="未选择" > 未选择</option>
				<option value="上学期" > 上学期</option>
				<option value="下学期" > 下学期</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width:auto;">是否核心课程</label>
	    <div class="layui-input-inline" >
		    <select name="Core" id="Core" lay-filter="style_select">
			    <option value="未选择" > 未选择</option>
				<option value="是" > 是</option>
				<option value="否" > 否</option>
	      	</select>
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
	    elem: '#SettingTime'
	    ,type: 'date'
	  });
	
	laydate.render({
	  elem: '#ApplyTime'
	  ,value: '20170910'
	  ,isInitValue: false //是否允许填充初始值，默认为 true
	});
	
	//文本域
	var index=layedit.build('detail',{
		hideTool:['image','face']
	});

	//数据上传
	function uploadData(){
		var Term=$("#Term").val()
		var Core=$("#Core").val()
		var data={
			'PlanId':[<.planId>],
			'CourseId':$("#CourseId").val(),
			'CourseName':$("#CourseName").val(),
			'Term':Term,
			'Core':Core
			};
			console.log(data)
			if(Term=="请选择"||Core=="请选择"){
				alert("不能为空")
			}else{
				//发布
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/jxjh/jxkc/allot_save",
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
			content: ['/v1/jxjh/jxkc/add'], //iframe的url，no代表不显示滚动条
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