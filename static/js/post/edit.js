var app = angular.module('itabirapp'); 
  app.controller('PostEditCtrl', function($scope, $http, $window) {

    this.delete = function (){
      var reg = /\/anuncio\/([0-9]+)\/editar/i;
      var postId = reg.exec($window.location.toString())[1];
      $http.get("/anuncio/" + postId + "/remover")
        .success(function(st) {
        alert("success")
        var res = st.Status;
        $scope.working = false;
        switch (res) {
          case "ok":
            alert('ok')
            $window.location.href = '/';
          break;
          case "err_usuario_inexiste":
            alert("err_usuario_inexiste");

          break;
          case "err_senha_invalida":
            alert("err_senha_invalida");
          break;
        }
      });
    }

    this.edit = function (){
        var obj, d, url;
        var reg = /\/anuncio\/([0-9]+)\/editar/i;
        var postId = reg.exec($window.location.toString())[1];
        d = this.post;
        obj = {Title: d.Title, Subtitle: d.Subtitle, Text: d.Text};

	    $http.post("/anuncio/" + postId + "/editar", obj).success(function(st) {
	        alert("success")
	        var res = st.Status;
	        $scope.working = false;
	        switch (res) {
	          case "ok":
	            alert('ok')
	          break;
	          case "err_usuario_inexiste":
	            alert("err_usuario_inexiste");

	          break;
	          case "err_senha_invalida":
	            alert("err_senha_invalida");
	          break;
	        }
	    });

      }
      /*
      $http.post('/login', {Email: d.Email, Password: d.Password}).
        */
    });
