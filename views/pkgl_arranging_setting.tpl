<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>排课参数设置</title>
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
	    <label class="layui-form-label" style="width: auto;">最大连排节数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="TogetherMax" autocomplete="off" class="layui-input">
	    </div>
		<div class="layui-form-mid layui-word-aux">*最小值为1</div>
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">文化课每天最大排课数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="CultureMax" autocomplete="off" class="layui-input">
	    </div>
		<div class="layui-form-mid layui-word-aux">*最小值为1</div>  
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">专业课每天最大排课数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="ProfessionalMax" autocomplete="off" class="layui-input">
	    </div>
		<div class="layui-form-mid layui-word-aux">*最小值为1</div> 
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">上午课下午不可排</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <select name="AmNotPm" id="AmNotPm" lay-filter="status_select">
			    <option value="选择" > 选择</option>
				<option value="是" > 是</option>
				<option value="否" >否</option>
	      </select>
	    </div>  
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">是否优先连排</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <select name="TogetherFirst" id="TogetherFirst" lay-filter="status_select">
			    <option value="选择" > 选择</option>
				<option value="是" > 是</option>
				<option value="否" >否</option>
	      </select>
	    </div>
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">文化课是否优先上午排课</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <select name="CultureAmFirst" id="CultureAmFirst" lay-filter="status_select">
			    <option value="选择" > 选择</option>
				<option value="是" > 是</option>
				<option value="否" >否</option>
	      </select>
	    </div>  
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">专业课是否优先下午排课</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <select name="ProfessionalPmFirst" id="ProfessionalPmFirst" lay-filter="status_select">
			    <option value="选择" > 选择</option>
				<option value="是" > 是</option>
				<option value="否" >否</option>
	      </select>
	    </div>  
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label" style="width: auto;">教师一周最大课时</label>
    <div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="TeacherMax" autocomplete="off" class="layui-input">
    </div>   
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label" style="width: auto;">同一班级同一科目每周最多节数</label>
    <div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="ClassCourseWeekMax" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label" style="width: auto;">每个班级每天最多课节数</label>
    <div class="layui-input-inline" style="width: 150px;">
      <input type="text" id="EachClassDayMax" autocomplete="off" class="layui-input">
    </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-inline" style="padding-left:50px;">
    	<button class="layui-btn" id="save">保存</button>
  	</div>
  </div>
</form>
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
	//自动加载
	$(function(){
				
	});	
 
	$('#save').on('click',function(){
		//alert("点击查询")
		var data={
			'Id':1,
			'TogetherMax':parseInt($("#AmNum").val()),
			'CultureMax':parseInt($("#PmNum").val()),
			'ProfessionalMax':parseInt($("#NightNum").val()),
			'TeacherMax':parseInt($("#AmNum").val()),
			'ClassCourseWeekMax':parseInt($("#PmNum").val()),
			'EachClassDayMax':parseInt($("#NightNum").val()),
			'AmNotPm':$("#AmNotPm").val(),
			'TogetherFirst':$("#TogetherFirst").val(),
			'CultureAmFirst':$("#CultureAmFirst").val(),
			'ProfessionalPmFirst':$("#ProfessionalPmFirst").val()
			};
				//发布
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/pkgl/setting/save",
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
		return false;
	})
});
</script>

</body>
</html>