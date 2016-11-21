var app = angular.module('itabirapp'); 
    app.controller('UserCtrlNew', function($scope, $http, $window) {
	    this.target = {disabled: true};
	    this.institution = {disabled: true};
	    this.address = {disabled: true};

	    this.newUser = function (){
		    //alert("olocomeu")
		    var obj, d;

		    d = this.target;
		    d2 = this.institution;
		    d3 = this.address;

		    obj = {
		      	Name: d.Name, Email: d.Email, Password: d.Password, Password2: d.Password2,
		      	Institution_Description: d2.Description, Institution_Tag: d2.Type,
		      	Addr_Street: d3.Street, Addr_Number: d3.Number, Addr_Complement: d3.Complement, 
		      	Addr_Neighborhood: d3.Neighborhood, Addr_City: d3.City
		    };

		    $http.post("/usuario/criar", obj).success(function(st) {
		        //alert("success")
		        var res = st.Status;
		        $scope.working = false;
		        switch (res) {
		          case "ok":
		            //alert('ok')
		            $window.location.href = '/';
		          break;
		        }
		    });

    	}


    });

$(function() {
 $("#imageUpload").change(function (){
 	$("#imageSend").removeAttr('disabled')
 });
});

function imageUpload() {
	//alert("xamooou")
	var data = new FormData();
	$.each($('#imageUpload')[0].files, function(i, file) {
	    data.append('datafile', file);
	});
	//alert("passou")

	$.ajax({
	    url: '/imagem/usuario/adicionar',
	    data: data,
	    cache: false,
	    contentType: false, //'multipart/form-data',
	    processData: false,
	    type: 'POST',
	    success: function(data){
	    	//alert("aehooooooooo")
	        //alert(data);
	    }
	});

}

