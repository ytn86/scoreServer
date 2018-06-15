function doRegister() {

	var json = {
		username: $('#input-username').val(),
		email: $('#input-email').val(),
		password: $('#input-password').val(),
		password2: $('#input-password2').val(),
		is_itf: $('#input-is_itf').val()
	};

	
	$.ajax({
		url: '/api/v1/register',
		headers: {
			'X-CSRF-Token': $('#csrf_token').val(),
		},
		data: JSON.stringify(json),
		contentType: 'application/json; charset=UTF-8',
		type: 'POST',
	}).done(function(data) {
		$('#msg').text(data.msg);
		if (data.msg=="success") {
			setTimeout(function() {
				location.href = '/login';
			}, 100);
		}
	}).fail(function(data) {
		$('#msg').text(data.responseJSON.msg);
	})
}


function doLogin() {

	var form = {username: $('#username').val(), password: $('#password').val()};
	var json = JSON.stringify(form);
	$.ajax({
		url: '/api/v1/login',
		headers: {
			'X-CSRF-Token': $('#csrf_token').val(),
		},
		data: json,
		contentType: 'application/json; charset=UTF-8',
		type: 'POST',
		
	}).done(function(data) {
		$('#msg').text(data.msg);
		if (data.msg=="success") {
			setTimeout(function() {
				location.href='/tasks';
			}, 1000);
		}
	}).fail(function(data) {
		$('#msg').text(data.responseJSON.msg);
	})
}


function submitFlag() {

	var form = {flag: $('#input-flag').val()};
	var taskid = $('#input-taskid').val();
	
	var json = JSON.stringify(form);

	$.ajax({
		url: '/api/v1/tasks/' + taskid + '/submit',
		headers: {
			'X-CSRF-Token': $('#csrf_token').val(),
		},
		data: json,
		contentType: 'application/json; charset=UTF-8',
		type: 'POST',
	}).done(function(data) {
		$('#msg').text(data.msg);
		if (data.msg=="congrats!") {
			$('#msg').text(data.msg);
			setTimeout(function() {
				$('#modal1').modal('close');
			}, 1000);
		}
	}).fail(function(data) {
		$('#msg').text(data.responseJSON.msg);
	})
}


function formToJson(selector) {

	var json = {};
	var formArray = selector.serializeArray();
	$.each(formArray, function() {
		if (!isNaN(this.value)) {
			json[this.name] = parseInt(this.value);
		} else if (this.value == 'false') {
			json[this.name] = false;
		} else if (this.value == 'true') {
			json[this.name] = true;
		} else {
			json[this.name] = this.value || '';
		}
	});
		   
	return JSON.stringify(json);
}

function dateToString(datetime) {

	var str;
	var date = new Date(datetime);
	str = date.getFullYear();
	str +=  '/';
	str += ('0'+(date.getMonth()+1)).substr(-2);
	str += '/';
	str += ('0'+date.getDate()).substr(-2);
	str += ' ';
	str += ('0'+date.getHours()).substr(-2);
	str += ':';
	str += ('0'+date.getMinutes()).substr(-2);;

	return str
}
