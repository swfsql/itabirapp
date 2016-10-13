$(document).ready(function() {
    $('#table').DataTable({
    	"order": [[ 0, "desc" ]],
    	"pageLength": 25,
        fixedHeader: {
            header: true,
            footer: true
        },
    	 "language": {
            "lengthMenu": "Mostrar _MENU_ usuários por página.",
            "zeroRecords": "Nenhum usuário encontrado.",
            "info": "Mostrando página _PAGE_ de _PAGES_.",
            "infoEmpty": "Nenhum usuário encontrado",
            "infoFiltered": "de um total de _MAX_ usuários.",
            "search" : "Busca:",
            "paginate" : {
            	"first" : "Primeira",
            	"last" : "Última",
            	"next" : "Próxima",
            	"previous" : "Anterior"
            },

        }
    });
} ); 
