// Generated by CoffeeScript 1.9.2
(function() {
  require.config({
    shim: {
      bscollapse: {
        deps: ['jquery']
      }
    },
    baseUrl: '/js/src',
    urlArgs: "bust=" + (new Date()).getTime(),
    paths: {
      domReady: 'vendor/requirejs-domready/2.0.1/domReady',
      headroom: 'vendor/headroom/0.7.0/headroom.min',
      jquery: 'vendor/jquery/2.1.4/jquery-2.1.4.min',
      bscollapse: 'vendor/bootstrap/3.3.5-collapse/bootstrap.min',
      react: 'vendor/react/0.13.3/react.min',
      jsdefines: 'lib/jsdefines'
    }
  });

  require(['jquery', 'react', 'jsdefines', 'domReady', 'headroom', 'bscollapse'], function($, React, jsdefines) {
    var neweventsource, newwebsocket, update, updates;
    updates = void 0;
    neweventsource = function(onmessage) {
      var conn, init, sendSearch;
      conn = null;
      sendSearch = function(search) {
        console.log('SEARCH', search);
        conn.close();
        return window.setTimeout(init, 1000);
      };
      init = function() {
        var again, statesel;
        conn = new EventSource('/index.sse' + location.search);
        conn.onopen = function() {
          $(window).bind('popstate', (function() {
            sendSearch(location.search);
          }));
        };
        statesel = 'table thead tr .header a.state';
        again = function(e) {
          $(statesel).unbind('click');
          if (!e.wasClean) {
            window.setTimeout(init, 5000);
          }
        };
        conn.onclose = function() {
          return console.log('sse closed (should recover)');
        };
        conn.onerror = function() {
          return console.log('sse errord (should recover)');
        };
        conn.onmessage = onmessage;
        $(statesel).click(function() {
          history.pushState({
            path: this.path
          }, '', this.href);
          sendSearch(this.search);
          return false;
        });
      };
      init();
      return {
        sendSearch: sendSearch,
        close: function() {
          return conn.close();
        }
      };
    };
    newwebsocket = function(onmessage) {
      var conn, init, sendSearch;
      conn = null;
      sendSearch = function(search) {
        console.log('Search', search);
        if ((conn == null) || conn.readyState === conn.CLOSING || conn.readyState === conn.CLOSED) {
          init();
        }
        if ((conn == null) || conn.readyState !== conn.OPEN) {
          console.log('Not connected, cannot send search', search);
          return;
        }
        return conn.send(JSON.stringify({
          Search: search
        }));
      };
      init = function() {
        var again, hostport, statesel;
        hostport = window.location.hostname + (location.port ? ':' + location.port : '');
        conn = new WebSocket('ws://' + hostport + '/index.ws');
        conn.onopen = function() {
          sendSearch(location.search);
          $(window).bind('popstate', (function() {
            sendSearch(location.search);
          }));
        };
        statesel = 'table thead tr .header a.state';
        again = function(e) {
          $(statesel).unbind('click');
          if (!e.wasClean) {
            window.setTimeout(init, 5000);
          }
        };
        conn.onclose = again;
        conn.onerror = again;
        conn.onmessage = onmessage;
        $(statesel).click(function() {
          history.pushState({
            path: this.path
          }, '', this.href);
          sendSearch(this.search);
          return false;
        });
      };
      init();
      return {
        sendSearch: sendSearch,
        close: function() {
          return conn.close();
        }
      };
    };
    this.IFCLASS = React.createClass({
      getInitialState: function() {
        return {
          Params: Data.Params,
          IFbytes: Data.IFbytes,
          IFerrors: Data.IFerrors,
          IFpackets: Data.IFpackets,
          ExpandableIF: Data.ExpandableIF,
          ExpandtextIF: Data.ExpandtextIF
        };
      },
      render: function() {
        var $if, Data;
        Data = this.state;
        return jsdefines.panelif.bind(this)(Data, (function() {
          var i, len, ref, ref1, ref2, results;
          ref2 = (ref = Data != null ? (ref1 = Data.IFpackets) != null ? ref1.List : void 0 : void 0) != null ? ref : [];
          results = [];
          for (i = 0, len = ref2.length; i < len; i++) {
            $if = ref2[i];
            results.push(jsdefines.ifpackets_rows(Data, $if));
          }
          return results;
        })(), (function() {
          var i, len, ref, ref1, ref2, results;
          ref2 = (ref = Data != null ? (ref1 = Data.IFerrors) != null ? ref1.List : void 0 : void 0) != null ? ref : [];
          results = [];
          for (i = 0, len = ref2.length; i < len; i++) {
            $if = ref2[i];
            results.push(jsdefines.iferrors_rows(Data, $if));
          }
          return results;
        })(), (function() {
          var i, len, ref, ref1, ref2, results;
          ref2 = (ref = Data != null ? (ref1 = Data.IFbytes) != null ? ref1.List : void 0 : void 0) != null ? ref : [];
          results = [];
          for (i = 0, len = ref2.length; i < len; i++) {
            $if = ref2[i];
            results.push(jsdefines.ifbytes_rows(Data, $if));
          }
          return results;
        })());
      },
      handleChange: function(e) {
        var href;
        href = '?' + e.target.name + '=' + e.target.value + '&' + location.search.substr(1);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      },
      handleClick: function(e) {
        var href;
        href = e.target.getAttribute('href');
        history.pushState({}, '', href);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      }
    });
    this.DFCLASS = React.createClass({
      getInitialState: function() {
        return {
          Params: Data.Params,
          DFbytes: Data.DFbytes,
          DFinodes: Data.DFinodes,
          ExpandableDF: Data.ExpandableDF,
          ExpandtextDF: Data.ExpandtextDF
        };
      },
      render: function() {
        var $disk, Data;
        Data = this.state;
        return jsdefines.paneldf.bind(this)(Data, (function() {
          var i, len, ref, ref1, ref2, results;
          ref2 = (ref = Data != null ? (ref1 = Data.DFinodes) != null ? ref1.List : void 0 : void 0) != null ? ref : [];
          results = [];
          for (i = 0, len = ref2.length; i < len; i++) {
            $disk = ref2[i];
            results.push(jsdefines.dfinodes_rows(Data, $disk));
          }
          return results;
        })(), (function() {
          var i, len, ref, ref1, ref2, results;
          ref2 = (ref = Data != null ? (ref1 = Data.DFbytes) != null ? ref1.List : void 0 : void 0) != null ? ref : [];
          results = [];
          for (i = 0, len = ref2.length; i < len; i++) {
            $disk = ref2[i];
            results.push(jsdefines.dfbytes_rows(Data, $disk));
          }
          return results;
        })());
      },
      handleChange: function(e) {
        var href;
        href = '?' + e.target.name + '=' + e.target.value + '&' + location.search.substr(1);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      },
      handleClick: function(e) {
        var href;
        href = e.target.getAttribute('href');
        history.pushState({}, '', href);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      }
    });
    this.LabelClassColorPercent = function(p) {
      if (p.length > 2) {
        return "label label-danger";
      }
      if (p.length > 1 && p[0] === '9') {
        return "label label-danger";
      }
      if (p.length > 1 && p[0] === '8') {
        return "label label-warning";
      }
      if (p.length > 1 && p[0] === '1') {
        return "label label-success";
      }
      if (p.length > 1) {
        return "label label-info";
      }
      return "label label-success";
    };
    this.MEMtableCLASS = React.createClass({
      getInitialState: function() {
        return {
          Params: Data.Params,
          MEM: Data.MEM
        };
      },
      render: function() {
        var $mem, Data;
        Data = this.state;
        return jsdefines.panelmem.bind(this)(Data, (function() {
          var i, len, ref, ref1, ref2, results;
          ref2 = (ref = Data != null ? (ref1 = Data.MEM) != null ? ref1.List : void 0 : void 0) != null ? ref : [];
          results = [];
          for (i = 0, len = ref2.length; i < len; i++) {
            $mem = ref2[i];
            results.push(jsdefines.mem_rows(Data, $mem));
          }
          return results;
        })());
      },
      handleChange: function(e) {
        var href;
        href = '?' + e.target.name + '=' + e.target.value + '&' + location.search.substr(1);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      },
      handleClick: function(e) {
        var href;
        href = e.target.getAttribute('href');
        history.pushState({}, '', href);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      }
    });
    this.CPUtableCLASS = React.createClass({
      getInitialState: function() {
        return {
          Params: Data.Params,
          CPU: Data.CPU
        };
      },
      render: function() {
        var $core, Data;
        Data = this.state;
        return jsdefines.panelcpu.bind(this)(Data, (function() {
          var i, len, ref, ref1, ref2, results;
          ref2 = (ref = Data != null ? (ref1 = Data.CPU) != null ? ref1.List : void 0 : void 0) != null ? ref : [];
          results = [];
          for (i = 0, len = ref2.length; i < len; i++) {
            $core = ref2[i];
            results.push(jsdefines.cpu_rows(Data, $core));
          }
          return results;
        })());
      },
      handleChange: function(e) {
        var href;
        href = '?' + e.target.name + '=' + e.target.value + '&' + location.search.substr(1);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      },
      handleClick: function(e) {
        var href;
        href = e.target.getAttribute('href');
        history.pushState({}, '', href);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      }
    });
    this.PStableCLASS = React.createClass({
      getInitialState: function() {
        return {
          Params: Data.Params,
          PStable: Data.PStable
        };
      },
      render: function() {
        var $proc, Data;
        Data = this.state;
        return jsdefines.panelps.bind(this)(Data, (function() {
          var i, len, ref, ref1, ref2, results;
          ref2 = (ref = Data != null ? (ref1 = Data.PStable) != null ? ref1.List : void 0 : void 0) != null ? ref : [];
          results = [];
          for (i = 0, len = ref2.length; i < len; i++) {
            $proc = ref2[i];
            results.push(jsdefines.ps_rows(Data, $proc));
          }
          return results;
        })());
      },
      handleChange: function(e) {
        var href;
        href = '?' + e.target.name + '=' + e.target.value + '&' + location.search.substr(1);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      },
      handleClick: function(e) {
        var href;
        href = e.target.getAttribute('href');
        history.pushState({}, '', href);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      }
    });
    this.VGtableCLASS = React.createClass({
      getInitialState: function() {
        return {
          Params: Data.Params,
          VagrantMachines: Data.VagrantMachines,
          VagrantError: Data.VagrantError,
          VagrantErrord: Data.VagrantErrord
        };
      },
      render: function() {
        var $mach, Data, rows;
        Data = this.state;
        if (((Data != null ? Data.VagrantErrord : void 0) != null) && Data.VagrantErrord) {
          rows = [jsdefines.vagrant_error.bind(this)(Data)];
        } else {
          rows = (function() {
            var i, len, ref, ref1, ref2, results;
            ref2 = (ref = Data != null ? (ref1 = Data.VagrantMachines) != null ? ref1.List : void 0 : void 0) != null ? ref : [];
            results = [];
            for (i = 0, len = ref2.length; i < len; i++) {
              $mach = ref2[i];
              results.push(jsdefines.vagrant_rows.bind(this)(Data, $mach));
            }
            return results;
          }).call(this);
        }
        return jsdefines.panelvg.bind(this)(Data, rows);
      },
      handleChange: function(e) {
        var href;
        href = '?' + e.target.name + '=' + e.target.value + '&' + location.search.substr(1);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      },
      handleClick: function(e) {
        var href;
        href = e.target.getAttribute('href');
        history.pushState({}, '', href);
        updates.sendSearch(href);
        e.stopPropagation();
        e.preventDefault();
        return void 0;
      }
    });
    this.NewTextCLASS = function(reduce) {
      return React.createClass({
        newstate: function(data) {
          var v;
          v = reduce(data);
          if (v != null) {
            return {
              Text: v
            };
          }
        },
        getInitialState: function() {
          return this.newstate(Data);
        },
        render: function() {
          return React.DOM.span(null, this.state.Text);
        }
      });
    };
    this.setState = function(obj, data) {
      var key;
      if (data != null) {
        for (key in data) {
          if (data[key] == null) {
            delete data[key];
          }
        }
        return obj.setState(data);
      }
    };
    update = function() {
      var cputable, dftable, hostname, iftable, ip, la, memtable, onmessage, pstable, uptime, vgtable;
      if ((typeof data !== "undefined" && data !== null ? data.IP : void 0) != null) {
        ip = React.render(React.createElement(NewTextCLASS(function(data) {
          return data != null ? data.IP : void 0;
        })), $('#ip').get(0));
      }
      hostname = React.render(React.createElement(NewTextCLASS(function(data) {
        return data != null ? data.Hostname : void 0;
      })), $('#hostname').get(0));
      uptime = React.render(React.createElement(NewTextCLASS(function(data) {
        return data != null ? data.Uptime : void 0;
      })), $('#uptime').get(0));
      la = React.render(React.createElement(NewTextCLASS(function(data) {
        return data != null ? data.LA : void 0;
      })), $('#la').get(0));
      memtable = React.render(React.createElement(MEMtableCLASS), document.getElementById('mem' + '-' + 'table'));
      pstable = React.render(React.createElement(PStableCLASS), document.getElementById('ps' + '-' + 'table'));
      dftable = React.render(React.createElement(DFCLASS), document.getElementById('df' + '-' + 'table'));
      cputable = React.render(React.createElement(CPUtableCLASS), document.getElementById('cpu' + '-' + 'table'));
      iftable = React.render(React.createElement(IFCLASS), document.getElementById('if' + '-' + 'table'));
      vgtable = React.render(React.createElement(VGtableCLASS), document.getElementById('vg' + '-' + 'table'));
      onmessage = function(event) {
        var data;
        data = JSON.parse(event.data);
        if (data == null) {
          return;
        }
        if ((data.Reload != null) && data.Reload) {
          window.setTimeout((function() {
            return location.reload(true);
          }), 5000);
          window.setTimeout(updates.close, 2000);
          console.log('in 5s: location.reload(true)');
          console.log('in 2s: updates.close()');
          return;
        }
        if (ip != null) {
          setState(ip, ip.newstate(data));
        }
        setState(hostname, hostname.newstate(data));
        setState(uptime, uptime.newstate(data));
        setState(la, la.newstate(data));
        setState(pstable, {
          Params: data.Params,
          PStable: data.PStable
        });
        setState(memtable, {
          Params: data.Params,
          MEM: data.MEM
        });
        setState(cputable, {
          Params: data.Params,
          CPU: data.CPU
        });
        setState(iftable, {
          Params: data.Params,
          IFbytes: data.IFbytes,
          IFerrors: data.IFerrors,
          IFpackets: data.IFpackets,
          ExpandableIF: data.ExpandableIF,
          ExpandtextIF: data.ExpandtextIF
        });
        setState(dftable, {
          Params: data.Params,
          DFbytes: data.DFbytes,
          DFinodes: data.DFinodes,
          ExpandableDF: data.ExpandableDF,
          ExpandtextDF: data.ExpandtextDF
        });
        setState(vgtable, {
          Params: data.Params,
          VagrantMachines: data.VagrantMachines,
          VagrantError: data.VagrantError,
          VagrantErrord: data.VagrantErrord
        });
        if (data.Location != null) {
          history.pushState({}, '', data.Location);
        }
      };
      updates = newwebsocket(onmessage);
    };
    require(['domReady', 'jquery', 'headroom'], function(domReady, $) {
      domReady(function() {
        var param;
        (new window.Headroom(document.querySelector('nav'), {
          offset: 20
        })).init();
        if (!((function() {
          var i, len, ref, results;
          ref = location.search.substr(1).split('&');
          results = [];
          for (i = 0, len = ref.length; i < len; i++) {
            param = ref[i];
            if (param.split('=')[0] === 'still') {
              results.push(42);
            }
          }
          return results;
        })()).length) {
          update();
        }
      });
    });
  });

}).call(this);
