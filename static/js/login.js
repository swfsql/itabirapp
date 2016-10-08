(function(){
  var app = angular.module('login', [ ]);  

  app.controller('LoginCtrl', function($scope, $http, $window) {
    $scope.login = {algo:"test"};

    this.sendLogin = function (){
      var d = this.login
      $http.post('/login', {Email: d.Email, Password: d.Password}).
        success(function(st) {
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
    };

  });

})();
