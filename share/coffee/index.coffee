require.config
  shim: {bscollapse: {deps: ['jquery']}} #, bootstrap: {deps: ['jquery']}
  baseUrl: '/js/src'
  urlArgs: "bust=" + (new Date()).getTime()
  paths:
    domReady:  'vendor/requirejs-domready/2.0.1/domReady'
    jquery:    'vendor/jquery/2.1.4/jquery.min'
    bscollapse:'vendor/bootstrap/3.3.5-collapse/bootstrap.min'
    react:     'vendor/react/0.13.3/react.min'
    jsdefines: 'lib/jsdefines'

# main require
require ['jquery', 'react', 'jsdefines', 'domReady', 'bscollapse'], ($, React, jsdefines) ->
  # domReady, bscollapse "required" for r.js only.
  updates = undefined # events source. set later
  neweventsource = (onmessage) ->
    conn = null
    sendSearch = (search) ->
      # conn = new EventSource('/index.sse' + search)
      # location.search = search
      console.log('SEARCH', search)
      conn.close() # if conn?
      window.setTimeout(init, 1000)
    init = () ->
      conn = new EventSource('/index.sse' + location.search)
      conn.onopen = () ->
        $(window).bind('popstate', (() ->
          sendSearch(location.search)
          return))
        return

      statesel = 'table thead tr .header a.state'
      again = (e) ->
        $(statesel).unbind('click')
        window.setTimeout(init, 5000) if !e.wasClean
        return

      conn.onclose   = () -> console.log('sse closed (should recover)')
      conn.onerror   = () -> console.log('sse errord (should recover)')
      conn.onmessage = onmessage

      $(statesel).click(() ->
        history.pushState({path: @path}, '', @href)
        sendSearch(@search)
        return false)
      return

    init()
    return {
      sendSearch: sendSearch
      close: () -> conn.close()
    }
  newwebsocket = (onmessage) ->
    conn = null
    sendSearch = (search) ->
      console.log 'Search', search
      # 0 conn.CONNECTING
      # 1 conn.OPEN
      # 2 conn.CLOSING
      # 3 conn.CLOSED
      if !conn? ||
         conn.readyState == conn.CLOSING ||
         conn.readyState == conn.CLOSED
        init()
      if !conn? ||
         conn.readyState != conn.OPEN
        console.log('Not connected, cannot send search', search)
        return
      return conn.send(JSON.stringify({Search: search}))
    init = () ->
      hostport = window.location.hostname +
        (if location.port then ':' + location.port else '')
      conn = new WebSocket('ws://' + hostport + '/index.ws')
      conn.onopen = () ->
        sendSearch(location.search)
        $(window).bind('popstate', (() ->
          sendSearch(location.search)
          return))
        return

      statesel = 'table thead tr .header a.state'
      again = (e) ->
        $(statesel).unbind('click')
        window.setTimeout(init, 5000) if !e.wasClean
        return

      conn.onclose   = again
      conn.onerror   = again
      conn.onmessage = onmessage

      $(statesel).click(() ->
        history.pushState({path: @path}, '', @href)
        sendSearch(@search)
        return false)
      return

    init()
    return {
      sendSearch: sendSearch
      close: () -> conn.close()
    }

  HandlerMixin =
    handleChange: (e) -> @handle(e, false, '?' + e.target.name + '=' + e.target.value + '&' + location.search.substr(1))
    handleClick: (e) ->
      href = e.target.getAttribute('href')
      href = $(e.target).parent().get(0).getAttribute('href') if !href?
      @handle(e, true, href)
    handle: (e, ps, href) ->
      history.pushState({}, '', href) if ps
      updates.sendSearch(href)
      e.stopPropagation() # preserves checkbox/radio
      e.preventDefault()  # checked/selected state
      return undefined

  @IFCLASS = React.createClass
    mixins: [HandlerMixin]
    getInitialState: () -> { # a global Data
      Params:       Data.Params
      IFbytes:      Data.IFbytes
      IFerrors:     Data.IFerrors
      IFpackets:    Data.IFpackets
      ExpandableIF: Data.ExpandableIF
      ExpandtextIF: Data.ExpandtextIF
    }
    render: () ->
      Data = @state
      return jsdefines.panelif.bind(this)(Data,
            (jsdefines.ifpackets_rows(Data, $if) for $if in Data?.IFpackets?.List ? []),
            (jsdefines.iferrors_rows(Data, $if) for $if in Data?.IFerrors?.List ? []),
            (jsdefines.ifbytes_rows(Data, $if) for $if in Data?.IFbytes?.List ? []))

  @DFCLASS = React.createClass
    mixins: [HandlerMixin]
    getInitialState: () -> { # a global Data
      Params:       Data.Params
      DFbytes:      Data.DFbytes
      DFinodes:     Data.DFinodes
      ExpandableDF: Data.ExpandableDF
      ExpandtextDF: Data.ExpandtextDF
    }
    render: () ->
      Data = @state
      return jsdefines.paneldf.bind(this)(Data,
             (jsdefines.dfinodes_rows(Data, $disk) for $disk in Data?.DFinodes?.List ? []),
             (jsdefines.dfbytes_rows(Data, $disk) for $disk in Data?.DFbytes?.List ? []))

  @MEMtableCLASS = React.createClass
    mixins: [HandlerMixin]
    getInitialState: () -> {
      Params: Data.Params,  # a global Data
      MEM:    Data.MEM,    # a global Data
    }
    render: () ->
      Data = @state
      return jsdefines.panelmem.bind(this)(Data, (jsdefines.mem_rows(Data, $mem
      ) for $mem in Data?.MEM?.List ? []))

  @CPUtableCLASS = React.createClass
    mixins: [HandlerMixin]
    getInitialState: () -> {
      Params: Data.Params, # a global Data
      CPU:    Data.CPU,    # a global Data
    }
    render: () ->
      Data = @state
      return jsdefines.panelcpu.bind(this)(Data, (jsdefines.cpu_rows(Data, $core
      ) for $core in Data?.CPU?.List ? []))

  @PStableCLASS = React.createClass
    mixins: [HandlerMixin]
    getInitialState: () -> {
      Params:  Data.Params, # a global Data
      PStable: Data.PStable # a global Data
    }
    render: () ->
      Data = @state
      return jsdefines.panelps.bind(this)(Data, (jsdefines.ps_rows(Data, $proc
      ) for $proc in Data?.PStable?.List ? []))

  @VGtableCLASS = React.createClass
    mixins: [HandlerMixin]
    getInitialState: () -> { # a global Data:
      Params:          Data.Params
      VagrantMachines: Data.VagrantMachines
      VagrantError:    Data.VagrantError
      VagrantErrord:   Data.VagrantErrord
    }
    render: () ->
      Data = @state
      if Data?.VagrantErrord? and Data.VagrantErrord
        rows = [jsdefines.vagrant_error.bind(this)(Data)]
      else
        rows = (jsdefines.vagrant_rows.bind(this)(Data, $mach
        ) for $mach in Data?.VagrantMachines?.List ? [])
      return jsdefines.panelvg.bind(this)(Data, rows)

  @NewTextCLASS = (reduce) -> React.createClass
    newstate: (data) ->
      v = reduce(data)
      return {Text: v} if v?
    getInitialState: () -> @newstate(Data) # a global Data
    render: () -> React.DOM.span(null, @state.Text)

  @setState = (obj, data) ->
    if data?
      delete data[key] for key of data when !data[key]?
      return obj.setState(data)

  update = () ->
    # coffeelint: disable=max_line_length
    ip = React.render(React.createElement(NewTextCLASS((data) -> data?.IP       )), $('#ip').get(0)) if data?.IP?
    hn = React.render(React.createElement(NewTextCLASS((data) -> data?.Hostname )), $('#hn').get(0))
    up = React.render(React.createElement(NewTextCLASS((data) -> data?.Uptime   )), $('#up').get(0))
    la = React.render(React.createElement(NewTextCLASS((data) -> data?.LA       )), $('#la').get(0))

    memtable  = React.render(React.createElement(MEMtableCLASS), document.getElementById('mem' +'-'+ 'table'))
    pstable   = React.render(React.createElement(PStableCLASS),  document.getElementById('ps'  +'-'+ 'table'))
    dftable   = React.render(React.createElement(DFCLASS),       document.getElementById('df'  +'-'+ 'table'))
    cputable  = React.render(React.createElement(CPUtableCLASS), document.getElementById('cpu' +'-'+ 'table'))
    iftable   = React.render(React.createElement(IFCLASS),       document.getElementById('if'  +'-'+ 'table'))
    vgtable   = React.render(React.createElement(VGtableCLASS),  document.getElementById('vg'  +'-'+ 'table'))
    # coffeelint: enable=max_line_length

    onmessage = (event) ->
      data = JSON.parse(event.data)
      return if !data?

      if data.Reload? and data.Reload
        window.setTimeout((() -> location.reload(true)), 5000)
        window.setTimeout(updates.close, 2000)
        console.log('in 5s: location.reload(true)')
        console.log('in 2s: updates.close()')
        return

      if data.Error?
        console.log 'Error', data.Error
        return

      setState(ip, ip.newstate(data)) if ip?
      setState(hn, hn.newstate(data))
      setState(up, up.newstate(data))
      setState(la, la.newstate(data))

      setState(pstable,  {Params: data.Params, PStable:  data.PStable})
      setState(memtable, {Params: data.Params, MEM: data.MEM})
      setState(cputable, {Params: data.Params, CPU: data.CPU})
      setState(iftable, {
        Params:       data.Params
        IFbytes:      data.IFbytes
        IFerrors:     data.IFerrors
        IFpackets:    data.IFpackets
        ExpandableIF: data.ExpandableIF
        ExpandtextIF: data.ExpandtextIF
      })
      setState(dftable, {
        Params:       data.Params
        DFbytes:      data.DFbytes
        DFinodes:     data.DFinodes
        ExpandableDF: data.ExpandableDF
        ExpandtextDF: data.ExpandtextDF
      })
      setState(vgtable, {
        Params:          data.Params,
        VagrantMachines: data.VagrantMachines,
        VagrantError:    data.VagrantError,
        VagrantErrord:   data.VagrantErrord
      })

      if data.Location?
        history.pushState({}, '', data.Location)

      return

    updates = newwebsocket(onmessage)
  # updates = neweventsource(onmessage)
    return # end of `update'

  require ['domReady', 'jquery'], (domReady, $) ->
    domReady () ->
      update() unless (42 for param in location.search.substr(1).split(
        '&') when (param.split('=')[0] == 'still')).length

# Local variables:
# coffee-tab-width: 2
# End:
