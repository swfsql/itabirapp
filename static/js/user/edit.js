(function(){
  var app = angular.module('itabirapp', [ ]);  

  app.controller('UserCtrl', function($scope, $http, $window) {
    this.target = {disabled: true};
    this.institution = {disabled: true};
    this.address = {disabled: true};

    this.toggle = function (){
      $http.get("/usuario/[[.Target.Id]]/editar/instituticao/autorizacao")
        .success(function(st) {
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
    
    this.send = function (form){
      var obj, d, url;
      var reg = /\/usuario\/([0-9]+)\/editar/i;
      var userId = reg.exec($window.location.toString())[1];
      switch (form) {
        case "target":
          d = this.user;
          this.user.disabled = true;
          url = "/usuario/" + userId + "/editar/usuario";
        break;
        case "institution":
          d = this.institution;
          alert("inst")
          this.institution.disabled = true;
          url = "/usuario/" + userId + "/editar/instituicao";
        break;
        case "address":
          d = this.address;
          this.address.disabled = true;
          obj = {Street: d.Street, Number: d.Number, Complement: d.Complement,
            Neighborhood: d.Neighborhood, City: d.City};
          url = "/usuario/" + userId + "/editar/endereco";
        break;
      }
      $http.post(url, obj).success(function(st) {
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

})();

