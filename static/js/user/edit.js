(function(){
  var app = angular.module('itabirapp', [ ]);  

  app.controller('UserCtrl', function($scope, $http, $window) {
    $scope.user = {disabled: true};
    $scope.institution = {disabled: true};
    $scope.address = {disabled: true};
    
    this.send = function (form){
      switch (form) {
        case "user":
          alert("user")
          $scope.user.disabled = true;
        break;
        case "institution":
          alert("inst")
          $scope.institution.disabled = true;
        break;
        case "address":
          $scope.address.disabled = true;
        break;
      }
      /*
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
        */
    };

  });

})();
