angular
.module('appServices')
.factory 'EmailTemplate', ['$resource', ($resource) ->
  $resource window.server_url + '/api/v1/email/template/:name', {id: '@name'},
    show:
      method: 'GET'
      responseType: 'json'
      transformResponse: (data) ->
        data?.template || data

    create:
      url: window.server_url + '/api/v1/email/template'
      method: 'POST'

    update:
      method: 'PUT'

    delete:
      method: 'DELETE'

    all:
      url: window.server_url + '/api/v1/email/templates'
      method: 'GET'
      responseType: 'json'
      transformResponse: (data) ->
        meta: data.meta
        items: data.templates

]