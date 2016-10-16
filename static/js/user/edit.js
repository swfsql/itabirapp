var app = angular.module('itabirapp'); 
  app.controller('UserEditCtrl', function($scope, $http, $window) {
    this.target = {disabled: true};
    this.institution = {disabled: true};
    this.address = {disabled: true};
    this.config = {disabled: true};


    this.delete = function (){
      var reg = /\/usuario\/([0-9]+)\/editar/i;
      var userId = reg.exec($window.location.toString())[1];
      $http.get("/usuario/" + userId + "/remover")
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

    this.toggle = function (){
      var reg = /\/usuario\/([0-9]+)\/editar/i;
      var userId = reg.exec($window.location.toString())[1];
      $http.get("/usuario/" + userId + "/editar/instituicao/autorizacao")
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
          d = this.target;
          this.target.disabled = true;
          obj = {Name: d.Name, Email: d.Email, Password: d.Password,
            Password2: d.Password2};
          url = "/usuario/" + userId + "/editar/usuario";
        break;
        case "institution":
          d = this.institution;
          alert("inst")
          this.institution.disabled = true;
          obj = {Description: d.Description};
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
  
$(function() {
 $("#imageUpload").change(function (){
  $("#imageSend").removeAttr('disabled')
 });
});

function imageUpload() {
  alert("xamooou")
  var data = new FormData();
  $.each($('#imageUpload')[0].files, function(i, file) {
      data.append('datafile', file);
  });
  alert("passou")

  $.ajax({
      url: '/imagem/usuario/adicionar',
      data: data,
      cache: false,
      contentType: false, //'multipart/form-data',
      processData: false,
      type: 'POST',
      success: function(data){
        alert("aehooooooooo")
          alert(data);
      }
  });

}
