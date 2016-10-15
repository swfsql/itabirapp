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
        obj = {Title: d.Title, Subtitle: d.Subtitle, Text: $window.simplemde.value()};


	    $http.post("/anuncio/" + postId + "/editar", obj).success(function(st) {
	        alert("success")
	        var res = st.Status;
	        $scope.working = false;
	        switch (res) {
	          case "ok":
	            alert('ok aki')
              $window.location.href = '/anuncio/' + postId;
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


var textArea = $("#textarea");
var simplemde = new SimpleMDE({ 
  element: textArea[0],
  autofocus: true,
      initialValue: textArea.text(),
      renderingConfig: {
        singleLineBreaks: true,
        codeSyntaxHighlighting: false,
    },
  spellChecker: false
  /*autosave: {
        enabled: true,
        uniqueId: "view/post/edit",
        delay: 6000, // each minute
    },*/ 
});
