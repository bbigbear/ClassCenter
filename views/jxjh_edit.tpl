<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>新建计划</title>
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
	<label class="layui-form-label">计划编号</label>
	<div class="layui-input-inline" >
		<input type="text" name="PlanId" id="PlanId" autocomplete="off" class="layui-input">	    	
	</div>
	<div class="layui-form-mid layui-word-aux">*修改时候只读</div>
  </div>
   <div class="layui-form-item">
	    <label class="layui-form-label">专业</label>
	    <div class="layui-input-inline" >
		    <select name="Major" id="Major" lay-filter="style_select">
			    <option value="专业名称" > 专业名称</option>
				<option value="计算机软件" > 计算机软件</option>
				<option value="物流服务与管理" > 物流服务与管理</option>
				<option value="市场营销" > 市场营销</option>
				<option value="计算机应用" > 计算机应用</option>
				<option value="电子商务" > 电子商务</option>
				<option value="信息管理" > 信息管理</option>
				<option value="电气自动化" > 电气自动化</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label">计划年级</label>
	    <div class="layui-input-inline" >
		    <select name="PlanGrade" id="PlanGrade" lay-filter="style_select">
			    <option value="2018" > 2018</option>
				<option value="2017" > 2017</option>
				<option value="2016" > 2016</option>
				<option value="2015" > 2015</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label">班级</label>
	    <div class="layui-input-inline" >
		    <select name="PlanClass" id="PlanClass" lay-filter="style_select">
			    <option value="1" > 1</option>
				<option value="2" > 2</option>
				<option value="3" > 3</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label">总学分要求</label>
	    <div class="layui-input-inline">
	      <input type="number" name="TotalCredits" id="TotalCredits" autocomplete="off" class="layui-input">
	    </div>
	<div class="layui-form-mid layui-word-aux">*如:170.5</div>
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label">建立年月</label>
	    <div class="layui-input-inline">
	      <input type="text" name="date" id="SettingTime" autocomplete="off" class="layui-input">
	    </div>
  </div>
 <div class="layui-form-item">
	    <label class="layui-form-label">使用学制</label>
	    <div class="layui-input-inline" >
			<input type="number" name="SchoolSystem" id="SchoolSystem" autocomplete="off" class="layui-input">
	    </div>		
	<div class="layui-form-mid layui-word-aux">*如:2.5</div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">培养目标</label>
    <div class="layui-input-block">
      <textarea class="layui-textarea" name="TrainTarget" id="TrainTarget"></textarea>
    </div>
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label">申请日期</label>
	    <div class="layui-input-inline">
	      <input type="text" name="date" id="ApplyTime" autocomplete="off" class="layui-input">
	    </div>
	<div class="layui-form-mid layui-word-aux">自动生成 不可修改</div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">申请说明</label>
    <div class="layui-input-block">
      <textarea class="layui-textarea" name="ApplyDescription" id="ApplyDescription"></textarea>
    </div>
  </div>
  <div class="layui-form-item">
    <div class="layui-input-block">
      <button class="layui-btn layui-btn-primary" id="save">保存</button>
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
	var id;
	$(function(){
		[<range .m>]
			id=[<.Id>]
			$("#PlanId").val([<.PlanId>])
			$("#Major").val([<.Major>])
			$("#PlanGrade").val([<.PlanGrade>])
			$("#PlanClass").val([<.PlanClass>])
			$("#TotalCredits").val([<.TotalCredits>])
			$("#SchoolSystem").val([<.SchoolSystem>])
			$("#TrainTarget").val([<.TrainTarget>])
			$("#ApplyDescription").val([<.ApplyDescription>])
			$("#Status").val([<.Status>])
			$("#SettingTime").val([<.SettingTime>])
			$("#ApplyTime").val([<.ApplyTime>])	
		[<end>]
	})
	
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
		var Major=$("#Major").val()
		var PlanGrade=$("#PlanGrade").val()
		var PlanClass=$("#PlanClass").val()
		var data={
			'Id':id,
			'PlanId':$("#PlanId").val(),
			'Major':Major,
			'PlanGrade':PlanGrade,
			'PlanClass':PlanClass,
			'TotalCredits':parseFloat($("#TotalCredits").val()),
			'SchoolSystem':parseFloat($("#SchoolSystem").val()),
			'TrainTarget':$("#TrainTarget").val(),
			'ApplyDescription':$("#ApplyDescription").val(),
			'Status':$("#Status").val(),
			'SettingTime':new Date($("#SettingTime").val()),
			'ApplyTime':new Date($("#ApplyTime").val())
			};
			console.log(data)
			if(Major=="请选择"||PlanGrade=="请选择"||PlanClass=="请选择"){
				alert("下拉列表不能为空")
			}else{
				//发布
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/jxjh/edit_action",
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
	
	$('#save').on('click',function(){
		uploadData()
		return false;
	});

	
	
});
</script>

</body>
</html>