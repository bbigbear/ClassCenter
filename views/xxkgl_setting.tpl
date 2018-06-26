<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>选课设置</title>
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
	    <label class="layui-form-label" style="width: auto;">每门选修课最多选课人数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
	    </div>
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">每门选修课班级最多选课人数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
	    </div> 
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">女生最大选课数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
	    </div>
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width: auto;">男生最大选课数</label>
	    <div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="PlanId" autocomplete="off" class="layui-input">
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
 
	$('#setting').on('click',function(){
		//alert("点击查询")
		var major=$("#Major").val()
		var grade=$("#PlanGrade").val()
		var classs=$("#PlanClass").val()
		var planId=$("#PlanId").val()
		table.reload('listReload',{
			where:{
				major:major,
				grade:grade,
				class:classs,
				planId:planId,
			}
		})		
		return false;
	})
});
</script>

</body>
</html>