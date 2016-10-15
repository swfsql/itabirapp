var app = angular.module('itabirapp'); 
    app.controller('PostNewCtrl', function($scope, $http, $window) {

	    this.newPost = function (){
		    alert("olocomeu")
		    var obj, d;

		    d = this.post;
		    var tags = ["tag_de_teste"];

		    obj = {Title: d.Title, Subtitle: d.Subtitle, Text: $window.simplemde.value(), 
		    	Tags: tags};

		    $http.post("/anuncio/criar", obj).success(function(st) {
		        alert("success")
		        var res = st.Status;
		        $scope.working = false;
		        switch (res) {
		          case "ok":
		            alert('ok ah vai mano')
		            $window.location.href = '/anuncio/' + st.PostId;
		          break;
		        }
		    });

    	}
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
