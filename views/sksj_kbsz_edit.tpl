<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>课表设置编辑</title>
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
	    <label class="layui-form-label">名称</label>
		<div class="layui-input-inline" style="width: 150px;">
	      <input type="text" id="KbszName" autocomplete="off" class="layui-input">
	    </div>
	    <!--<div class="layui-input-inline" >
		    <select name="KbszName" id="KbszName" lay-filter="style_select">
			    <option value="未选择" > 未选择</option>
				<option value="早自习" > 早自习</option>
				<option value="第一节课" > 第一节课</option>
				<option value="第二节课" > 第二节课</option>
				<option value="第三节课" > 第三节课</option>
				<option value="第四节课" > 第四节课</option>
				<option value="第五节课" > 第五节课</option>
				<option value="第六节课" > 第六节课</option>
				<option value="晚自习" > 晚自习</option>
	      	</select>
		</div>-->		
  </div>
  <div class="layui-form-item">
	<label class="layui-form-label">开始时间</label>
	<div class="layui-input-inline" >
		<input type="text" name="date" id="StartTime" autocomplete="off" class="layui-input">	    	
	</div>
  </div> 
  <div class="layui-form-item">
	<label class="layui-form-label">结束时间</label>
	<div class="layui-input-inline" >
		<input type="text" name="date" id="EndTime" autocomplete="off" class="layui-input">	    	
	</div>
  </div> 
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width:auto;">周一是否启动</label>
	    <div class="layui-input-inline" >
		    <select name="Mon" id="Mon" lay-filter="style_select">
				<option value="" > 未选择</option>
				<option value="启动" > 启动</option>
				<option value="未启动" > 未启动</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width:auto;">周二是否启动</label>
	    <div class="layui-input-inline" >
		    <select name="Tue" id="Tue" lay-filter="style_select">
				<option value="" > 未选择</option>
				<option value="启动" > 启动</option>
				<option value="未启动" > 未启动</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width:auto;">周三是否启动</label>
	    <div class="layui-input-inline" >
		    <select name="Wed" id="Wed" lay-filter="style_select">
				<option value="" > 未选择</option>
				<option value="启动" > 启动</option>
				<option value="未启动" > 未启动</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width:auto;">周四是否启动</label>
	    <div class="layui-input-inline" >
		    <select name="Thu" id="Thu" lay-filter="style_select">
				<option value="" > 未选择</option>
				<option value="启动" > 启动</option>
				<option value="未启动" > 未启动</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width:auto;">周五是否启动</label>
	    <div class="layui-input-inline" >
		    <select name="Fri" id="Fri" lay-filter="style_select">
				<option value="" > 未选择</option>
				<option value="启动" > 启动</option>
				<option value="未启动" > 未启动</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width:auto;">周六是否启动</label>
	    <div class="layui-input-inline" >
		    <select name="Sat" id="Sat" lay-filter="style_select">
				<option value="未选择" > 未选择</option>
				<option value="启动" > 启动</option>
				<option value="未启动" > 未启动</option>
	      	</select>
		</div>		
  </div>
  <div class="layui-form-item">
	    <label class="layui-form-label" style="width:auto;">周日是否启动</label>
	    <div class="layui-input-inline" >
		    <select name="Sun" id="Sun" lay-filter="style_select">
				<option value="" > 未选择</option>
				<option value="启动" > 启动</option>
				<option value="未启动" > 未启动</option>
	      	</select>
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
	var style;
	$(function(){
		[<range .m>]
			id=[<.Id>]
			$("#KbszName").val([<.Name>])
			$("#StartTime").val([<.StartTime>])
			$("#EndTime").val([<.EndTime>])
			$("#StartTime").val([<.StartTime>])
			$("#Mon").val([<.Mon>])
			$("#Tue").val([<.Tue>])
			$("#Wed").val([<.Wed>])
			$("#Thu").val([<.Thu>])
			$("#Fri").val([<.Fri>])
			$("#Sat").val([<.Sat>])
			$("#Sun").val([<.Sun>])
			style=[<.Style>]
		[<end>]
	})
	laydate.render({
	    elem: '#StartTime'
	    ,type: 'time'
	  });
	
	laydate.render({
	    elem: '#EndTime'
	    ,type: 'time'
	  });

	//数据上传
	function uploadData(){
		var data={
				'Id':id,
				'Name':$("#KbszName").val(),
				'Style':style,
				'StartTime':$("#StartTime").val(),
				'EndTime':$("#EndTime").val(),
				'Mon':$("#Mon").val(),
				'Tue':$("#Tue").val(),
				'Wed':$("#Wed").val(),
				'Thu':$("#Thu").val(),
				'Fri':$("#Fri").val(),
				'Sat':$("#Sat").val(),
				'Sun':$("#Sun").val()			
			};

				//发布
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/sksj/edit_action",
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
	
	$('#save').on('click',function(){
		uploadData()
		return false;
	});
	
});
</script>

</body>
</html>