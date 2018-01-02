<!DOCTYPE html>
<html>
<head>
	{{template "header"}}
	<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
	<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
</head>
<body>

<div class="navbar navbar-default" >
	<div class="container">
	<div class="navbar-header">
	    <a class="navbar-brand" href="#">在线选课</a>
	</div>
	</div>
</div>
<div class="container">	
		<div class="panel panel-default">
		    <div class="panel-body">
			<form class="form-inline" role="form">
		        <div class="form-group">				
					<label>学年学期</label>			
					<select class="form-control" >
						<option>2017-2018学年春季学期</option>
						<option>2017-2018学年秋季学期</option>
					</select>				
				</div>
				<div class="form-group" style="padding-top:10px">				
					<label>培养层次</label>			
					<select class="form-control">
						<option>博士</option>
						<option>本科</option>
					</select>				
					<label>年纪</label>				
					<select class="form-control">
						<option>2015</option>
						<option>2016</option>
					</select>				
					<label>院(系)</label>	
					<select class="form-control">
						<option>计算机科学与技术系</option>
						<option>音乐学院</option>
					</select>
					<label>专业</label>
					<select class="form-control">
						<option>计算机科学与技术</option>
						<option>软件工程</option>
					</select>		
				</div>
				<div class="form-group" style="padding-top:10px">				
					<label>课程类别</label>			
					<select class="form-control">
						<option></option>
						<option></option>
					</select>							
					<select class="form-control">
						<option></option>
						<option></option>
					</select>				
					<select class="form-control">
						<option></option>
						<option></option>
					</select>
					<label>课程</label>
					<input class="form-control">
					<label>任课老师</label>
					<input class="form-control">		
				</div>
				<div class="form-group" style="padding-top:10px">				
					<div class="radio">
					    <label>
					        <input type="radio" name="optionsRadios" id="optionsRadios1" value="option1" checked> 限本院系开设课程
					    </label>
						<label>
					        <input type="radio" name="optionsRadios" id="optionsRadios2" value="option2" > 其他院系开设课程
					    </label>
						<label>
					        <input type="radio" name="optionsRadios" id="optionsRadios3" value="option3" > 所有院系开设课程
					    </label>
					</div>				
				</div>
			</form>				 	
		    </div>
		</div>
		<div class="table-responsive">
					<table class="table table-bordered">
						<caption>课表</caption>
						<thead>
							<tr >
								<th>序号</th>
								<th>课程</th>
								<th>学分</th>
								<th>总学时</th>
								<th>课程类别</th>
								<th>开课单位</th>
								<th>上课班号</th>
								<th>限选人数</th>
								<th>任课老师</th>
							</tr>
						</thead>
						<tbody>							
						</tbody>
				</table>
		</div> 
</div>
</body>
</html>
