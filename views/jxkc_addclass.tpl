<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>新建教学课程</title>
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
		    <select name="trainChannel" id="trainChannel" lay-filter="style_select">
				<option value="必修课程" > 必修课程</option>
				<option value="选修课程" > 选修课程</option>
	      	</select>
		</div>
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">课程类别</label>
	    <div class="layui-input-block" >
		    <select name="trainChannel" id="trainChannel" lay-filter="style_select">
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
			<input type="text" name="department" id="department" autocomplete="off" class="layui-input">
	    </div>
		<div class="layui-form-mid layui-word-aux">*修改时候只读</div>
	</div>
  </div>
  <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">课程名称</label>
	    <div class="layui-input-block" >
		<input type="text" name="department" id="department" autocomplete="off" class="layui-input">
	    </div>
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">课程英文名</label>
	    <div class="layui-input-block" >
		<input type="text" name="leader" id="leader" autocomplete="off" class="layui-input">
	    </div>
	</div>
  </div>
  <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">课程别名</label>
	    <div class="layui-input-block" >
		<input type="text" name="joinNum" id="joinNum" autocomplete="off" class="layui-input" lay-verify="required|number">
	    </div>		
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">课程费用</label>
	    <div class="layui-input-block" >
		<input type="text" name="trainPlace" id="trainPlace" autocomplete="off" class="layui-input">
	    </div>
	</div>
  </div>
 <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">开课机构</label>
	    <div class="layui-input-block" >
		    <select name="trainChannel" id="trainChannel" lay-filter="style_select">
			    <option value="请选择" > 请选择</option>
				<option value="学校机构" > 学校机构</option>
	      	</select>
		</div>		
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">课程系数</label>
	    <div class="layui-input-block" >
		<input type="text" name="trainContract" id="trainContract" autocomplete="off" class="layui-input">
	    </div>
	</div>
  </div>
 <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">考试形式</label>
	    <div class="layui-input-block" >
		    <select name="trainChannel" id="trainChannel" lay-filter="style_select">
			    <option value="请选择" > 请选择</option>
				<option value="考试" > 考试</option>
				<option value="考查" > 考查</option>
	      	</select>
		</div>		
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">授课方式</label>
	    <div class="layui-input-block" >
		    <select name="trainChannel" id="trainChannel" lay-filter="style_select">
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
	      <input type="text" name="date" id="date1" autocomplete="off" class="layui-input">
	    </div>
	</div>
	<div class="layui-inline">
		<label class="layui-form-label">理论学时</label>
		<div class="layui-input-inline">
	      <input type="text" name="date" id="date2" autocomplete="off" class="layui-input">
	    </div>
 	</div>
  </div>
 <div class="layui-form-item">
	<div class="layui-inline">
	    <label class="layui-form-label">实践学时</label>
	    <div class="layui-input-block" >
		<input type="text" name="budget" id="budget" autocomplete="off" class="layui-input" lay-verify="required|number">
	    </div>		
	</div>
	<div class="layui-inline">
	    <label class="layui-form-label">其他学时</label>
	    <div class="layui-input-block" >
		<input type="text" name="approver" id="approver" autocomplete="off" class="layui-input">
	    </div>
	</div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">参考书目</label>
    <div class="layui-input-block" style="width: 500px;">
	<input type="text" name="joinDepartment" id="joinDepartment" autocomplete="off" class="layui-input">
    </div>
	<div class="layui-form-mid layui-word-aux">多本用,隔开</div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">学分</label>
    <div class="layui-input-block" style="width: 500px;">
	<input type="text" name="joinPeople" id="joinPeople" autocomplete="off" class="layui-input">
    </div>
	<div class="layui-form-mid layui-word-aux">*如:2.5</div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">课程介绍</label>
    <div class="layui-input-block">
      <textarea class="layui-textarea" name="trainInfo" id="trainInfo"></textarea>
    </div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">排课用优化信息:</label>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">不能排课节次:</label>
   	<div class="layui-input-block" style="width: 500px;">
		<input type="text" name="joinPeople" id="joinPeople" autocomplete="off" class="layui-input">
    </div>
	<div class="layui-form-mid layui-word-aux">多个节次用逗号隔开 例1,7,8 节次<=10</div>
  </div>
  <div class="layui-form-item">
    <label class="layui-form-label">优先排课次数:</label>
    <div class="layui-input-block" style="width: 500px;">
		<input type="text" name="joinPeople" id="joinPeople" autocomplete="off" class="layui-input">
    </div>
	<div class="layui-form-mid layui-word-aux">多个节次用逗号隔开 例2,3 节次<=10</div>
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
	    elem: '#day'
		,type: 'datetime'
	  });
	
	 laydate.render({
	    elem: '#date1'
	    //,type: 'date'
		,type: 'datetime'
	  });
	
	 laydate.render({
	    elem: '#date2'
	    ,type: 'datetime'
	  });
	
	//文本域
	var index=layedit.build('detail',{
		hideTool:['image','face']
	});
		
   
	  //多文件列表示例
	  var path=""
	  var demoListView = $('#demoList')
	  ,uploadListIns = upload.render({
	    elem: '#testList'
	    ,url: '/v1/put_file'
	    ,accept: 'file'
	    ,multiple: true
	    ,auto: false
	    ,bindAction: '#save'
	    ,choose: function(obj){   
	      var files = this.files = obj.pushFile(); //将每次选择的文件追加到文件队列
	      //读取本地文件
	      obj.preview(function(index, file, result){
	        var tr = $(['<tr id="upload-'+ index +'">'
	          ,'<td>'+ file.name +'</td>'
	          ,'<td>'+ (file.size/1014).toFixed(1) +'kb</td>'
	          ,'<td>等待上传</td>'
	          ,'<td>'
	            ,'<button class="layui-btn layui-btn-xs demo-reload layui-hide">重传</button>'
	            ,'<button class="layui-btn layui-btn-xs layui-btn-danger demo-delete">删除</button>'
	          ,'</td>'
	        ,'</tr>'].join(''));
	        
	        //单个重传
	        tr.find('.demo-reload').on('click', function(){
	          obj.upload(index, file);
	        });
	        
	        //删除
	        tr.find('.demo-delete').on('click', function(){
	          delete files[index]; //删除对应的文件
	          tr.remove();
	          uploadListIns.config.elem.next()[0].value = ''; //清空 input file 值，以免删除后出现同名文件不可选
	        });
	        
	        demoListView.append(tr);
	      });
	    }
	    ,done: function(res, index, upload){
	      if(res.code == 200){ //上传成功
			path=path+res.data.src+','
	        var tr = demoListView.find('tr#upload-'+ index)
	        ,tds = tr.children();
	        tds.eq(2).html('<span style="color: #5FB878;">上传成功</span>');
	        tds.eq(3).html(''); //清空操作
	        return delete this.files[index]; //删除文件队列已经上传成功的文件
	      }
	      this.error(index, upload);
	    }
	    ,error: function(index, upload){
	      var tr = demoListView.find('tr#upload-'+ index)
	      ,tds = tr.children();
	      tds.eq(2).html('<span style="color: #FF5722;">上传失败</span>');
	      tds.eq(3).find('.demo-reload').removeClass('layui-hide'); //显示重传
	    }
		,allDone: function(obj){
	      	//alert(path_src)
			console.log(obj)
			//post json
			uploadData();						
	    }
	  });
		
	//数据上传
	function uploadData(){
		var trainChannel=$("#trainChannel").val()
		var trainForm=$("#trainForm").val()
		var data={
			'number':$("#number").val(),
			'planName':$("#planName").val(),
			'trainChannel':trainChannel,
			'trainForm':trainForm,
			'department':$("#department").val(),
			'leader':$("#leader").val(),
			'joinNum':parseInt($("#joinNum").val()),
			'trainPlace':$("#trainPlace").val(),
			'trainName':$("#trainName").val(),
			'trainContract':$("#trainContract").val(),
			'trainCname':$("#trainCname").val(),
			'trainTotalTime':parseInt($("#totalTime").val()),			
			'budget':parseInt($("#budget").val()),
			'approver':$("#approver").val(),
			'joinDepartment':$("#joinDepartment").val(),
			'joinPeople':$("#joinPeople").val(),
			'trainInfo':$("#trainInfo").val(),
			'trainContractInfo':$("#trainContractInfo").val(),
			'trainRequire':$("#trainRequire").val(),
			'trainExplain':$("#trainExplain").val(),
			'remark':$("#remark").val(),
			'filePath':path,
			'trainDetail':$("#trainDetail").val(),
			'trainStartTime':$("#date1").val(),
			'trainEndTime':$("#date2").val()
			};
			console.log(data)
			if(trainChannel=="请选择"||trainForm=="请选择"){
				alert("请选择培训通道或者培训形式")
			}else if($("#date1").val()==""||$("#date2").val()==""){
				alert("时间不能为空")
			}else{
				//发布
				$.ajax({
					type:"POST",
					contentType:"application/json;charset=utf-8",
					url:"/v1/train/plan/add_action",
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
	
	//保存
	$('#save').on('click',function(){
		//console.log($("#demoList tr").length)			
		var len = $("#demoList tr").length
		if (len==0){
			uploadData()
		}
		return false;
	});
});
</script>

</body>
</html>