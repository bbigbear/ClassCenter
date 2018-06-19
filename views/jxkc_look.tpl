<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>查看教学课程</title>
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
	<div class="layui-inline">
	    <label class="layui-form-label">选修类别</label>
	    <div class="layui-input-block" >
		    <select name="ElectiveCategory" id="ElectiveCategory" lay-filter="style_select">
				<option value="必修课程" > 必修课程</option>
				<option value="选修课程" > 选修课程</option>
	      	</select>
		</div>
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">课程类别</label>
	    <div class="layui-input-block" >
		    <select name="CourseCategory" id="CourseCategory" lay-filter="style_select">
				<option value="文化课" > 文化课</option>
				<option value="专业课" > 专业课</option>
	      	</select>
		</div>
	</div>
  </div>
  <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">课程号</label>
		<div class="layui-input-block" >
			<input type="text" name="CourseId" id="CourseId" autocomplete="off" class="layui-input">
	    	<div class="layui-form-mid layui-word-aux">*修改时候只读</div>
		</div>
	</div>
  </div>
  <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">课程名称</label>
	    <div class="layui-input-block" >
		<input type="text" name="CourseName" id="CourseName" autocomplete="off" class="layui-input">
	    </div>
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">课程英文名</label>
	    <div class="layui-input-block" >
		<input type="text" name="CourseEnName" id="CourseEnName" autocomplete="off" class="layui-input">
	    </div>
	</div>
  </div>
  <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">课程别名</label>
	    <div class="layui-input-block" >
		<input type="text" name="CourseAlias" id="CourseAlias" autocomplete="off" class="layui-input" lay-verify="required|number">
	    </div>		
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">课程费用</label>
	    <div class="layui-input-block" >
		<input type="text" name="CourseCost" id="CourseCost" autocomplete="off" class="layui-input">
	    </div>
	</div>
  </div>
 <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">开课机构</label>
	    <div class="layui-input-block" >
		    <select name="OpeningInstitution" id="OpeningInstitution" lay-filter="style_select">
			    <option value="请选择" > 请选择</option>
				<option value="学校机构" > 学校机构</option>
	      	</select>
		</div>		
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">课程系数</label>
	    <div class="layui-input-block" >
		<input type="text" name="OpeningCoefficient" id="OpeningCoefficient" autocomplete="off" class="layui-input">
	    </div>
	</div>
  </div>
 <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">考试形式</label>
	    <div class="layui-input-block" >
		    <select name="ExaminationForm" id="ExaminationForm" lay-filter="style_select">
			    <option value="请选择" > 请选择</option>
				<option value="考试" > 考试</option>
				<option value="考查" > 考查</option>
	      	</select>
		</div>		
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">授课方式</label>
	    <div class="layui-input-block" >
		    <select name="TeachingMethods" id="TeachingMethods" lay-filter="style_select">
			    <option value="请选择" > 请选择</option>
				<option value="面授讲课" > 面授讲课</option>
				<option value="辅导" > 辅导</option>
				<option value="录像讲课" > 录像讲课</option>
				<option value="网络教学" > 网络教学</option>
				<option value="实验" > 实验</option>
				<option value="实习" > 实习</option>
				<option value="其他" > 其他</option>
	      	</select>
		</div>
	</div>
  </div>
  <div class="layui-form-item">
    <div class="layui-inline">
	    <label class="layui-form-label">总学时</label>
	    <div class="layui-input-inline">
	      <input type="number" name="TotalHours" id="TotalHours" autocomplete="off" class="layui-input">
	    </div>
	</div>
	<div class="layui-inline">
		<label class="layui-form-label">理论学时</label>
		<div class="layui-input-inline">
	      <input type="number" name="TheoreticalHours" id="TheoreticalHours" autocomplete="off" class="layui-input">
	    </div>
 	</div>
  </div>
 <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">实践学时</label>
	    <div class="layui-input-block" >
		<input type="number" name="PracticeHours" id="PracticeHours" autocomplete="off" class="layui-input">
	    </div>		
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">其他学时</label>
	    <div class="layui-input-block" >
		<input type="number" name="OtherHours" id="OtherHours" autocomplete="off" class="layui-input">
	    </div>
	</div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">参考书目</label>
    <div class="layui-input-block" style="width: 500px;">
		<input type="text" name="Bibliography" id="Bibliography" autocomplete="off" class="layui-input">
    	<div class="layui-form-mid layui-word-aux">多本用,隔开</div>
	</div>
	
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">学分</label>
    <div class="layui-input-block" style="width: 500px;">
		<input type="number" name="Credit" id="Credit" autocomplete="off" class="layui-input">
    	<div class="layui-form-mid layui-word-aux">*如:2.5</div>
	</div>	
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">课程介绍</label>
    <div class="layui-input-block">
      <textarea class="layui-textarea" name="CourseIntroduction" id="CourseIntroduction"></textarea>
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label" style="width:110px;color:#01AAED;">排课用优化信息</label>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">不能排课节次</label>
   	<div class="layui-input-block" style="width: 500px;">
		<input type="text" name="NoCourseNumber" id="NoCourseNumber" autocomplete="off" class="layui-input">
   		<div class="layui-form-mid layui-word-aux">多个节次用逗号隔开 例1,7,8 节次<=10</div>	
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">优先排课次数</label>
    <div class="layui-input-block" style="width: 500px;">
		<input type="text" name="CourseNumber" id="CourseNumber" autocomplete="off" class="layui-input">
    	<div class="layui-form-mid layui-word-aux">多个节次用逗号隔开 例2,3 节次<=10</div>
	</div>	
  </div>
  <div class="layui-form-item">
    <div class="layui-input-block">
<!--      <button class="layui-btn layui-btn-primary" id="add">添加</button>-->
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
			$("#ElectiveCategory").val([<.ElectiveCategory>])
			$("#CourseCategory").val([<.CourseCategory>])
			$("#CourseId").val([<.CourseId>])
			$("#CourseName").val([<.CourseName>])
			$("#CourseEnName").val([<.CourseEnName>])
			$("#CourseAlias").val([<.CourseAlias>])
			$("#CourseCost").val([<.CourseCost>])
			$("#OpeningInstitution").val([<.OpeningInstitution>])
			$("#OpeningCoefficient").val([<.OpeningCoefficient>])
			$("#ExaminationForm").val([<.ExaminationForm>])
			$("#TeachingMethods").val([<.TeachingMethods>])
			$("#TotalHours").val([<.TotalHours>])	
			$("#TheoreticalHours").val([<.TheoreticalHours>])
			$("#PracticeHours").val([<.PracticeHours>])
			$("#OtherHours").val([<.OtherHours>])
			$("#Bibliography").val([<.Bibliography>])
			$("#Credit").val([<.Credit>])
			$("#CourseIntroduction").val([<.CourseIntroduction>])
			$("#CourseNumber").val([<.CourseNumber>])
			$("#NoCourseNumber").val([<.NoCourseNumber>])		
		[<end>]
	})
	//文本域
	var index=layedit.build('detail',{
		hideTool:['image','face']
	});

	//数据上传
	function uploadData(){
		var ExaminationForm=$("#ExaminationForm").val()
		var TeachingMethods=$("#TeachingMethods").val()
		var OpeningInstitution=$("#OpeningInstitution").val()
		var data={
			'Id':id,
			'ElectiveCategory':$("#ElectiveCategory").val(),
			'CourseCategory':$("#CourseCategory").val(),
			'CourseId':$("#CourseId").val(),
			'CourseName':$("#CourseName").val(),
			'CourseEnName':$("#CourseEnName").val(),
			'CourseAlias':$("#CourseAlias").val(),
			'CourseCost':$("#CourseCost").val(),
			'OpeningInstitution':OpeningInstitution,
			'OpeningCoefficient':$("#OpeningCoefficient").val(),
			'ExaminationForm':ExaminationForm,
			'TeachingMethods':TeachingMethods,
			'TotalHours':parseFloat($("#TotalHours").val()),			
			'TheoreticalHours':parseFloat($("#TheoreticalHours").val()),
			'PracticeHours':parseFloat($("#PracticeHours").val()),
			'OtherHours':parseFloat($("#OtherHours").val()),
			'Bibliography':$("#Bibliography").val(),
			'Credit':parseFloat($("#Credit").val()),
			'CourseIntroduction':$("#CourseIntroduction").val(),
			'CourseNumber':$("#CourseNumber").val(),
			'NoCourseNumber':$("#NoCourseNumber").val()
			};
			console.log(data)
			if(ExaminationForm=="请选择"||TeachingMethods=="请选择"||OpeningInstitution=="请选择"){
				alert("下拉列表不能为空")
			}else{
				//发布
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/jxkc/edit_action",
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