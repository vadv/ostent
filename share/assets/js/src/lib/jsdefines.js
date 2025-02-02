"use strict";

define(function (require) {
  var React = require("react");
  return {
    mem_rows: function mem_rows(Data, $mem) {
      return React.createElement(
        "tr",
        { key: "mem-rowby-kind-" + $mem.Kind
        },
        React.createElement(
          "td",
          null,
          $mem.Kind
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $mem.Free
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          React.createElement(
            "span",
            { className: "label", "data-usepercent": $mem.UsePercent
            },
            $mem.UsePercent,
            "%"
          ),
          " ",
          $mem.Used
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $mem.Total
        )
      );
    },
    panelmem: function panelmem(Data, rows) {
      return React.createElement(
        "div",
        null,
        React.createElement(
          "div",
          null,
          React.createElement(
            "a",
            { href: Data.Params.BOOL.configmem.Href, onClick: this.handleClick, className: "btn-block"
            },
            React.createElement(
              "span",
              { className: Data.Params.BOOL.configmem.Value ? "h4 bg-info" : "h4"
              },
              "Memory"
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { id: "memconfig", className: Data.Params.BOOL.configmem.Value ? "config-margintop" : "config-margintop collapse-hidden"
            },
            React.createElement(
              "form",
              { className: "form-inline", action: "/form/" + Data.Params.Query
              },
              React.createElement("input", { className: "hidden-submit", type: "submit"
              }),
              React.createElement(
                "div",
                { className: "btn-toolbar"
                },
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "a",
                    { className: Data.Params.BOOL.hidemem.Value ? "btn btn-default active" : "btn btn-default", href: Data.Params.BOOL.hidemem.Href, onClick: this.handleClick
                    },
                    "Hidden"
                  ),
                  React.createElement(
                    "a",
                    { className: Data.Params.BOOL.hideswap.Value ? "btn btn-default active" : "btn btn-default", href: Data.Params.BOOL.hideswap.Href, onClick: this.handleClick, disabled: Data.Params.BOOL.hidemem.Value ? "disabled" : ""
                    },
                    "Hide swap"
                  )
                ),
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "div",
                    { className: "input-group input-group-sm refresh-group" + (Data.Params.PERIOD.refreshmem.InputErrd ? " has-warning" : "")
                    },
                    React.createElement(
                      "span",
                      { className: "input-group-addon"
                      },
                      "Refresh"
                    ),
                    React.createElement("input", { className: "form-control refresh-input width-fourem", type: "text", placeholder: Data.Params.PERIOD.refreshmem.Placeholder, name: "refreshmem", value: Data.Params.PERIOD.refreshmem.Input, onChange: this.handleChange
                    })
                  )
                )
              )
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { className: Data.Params.BOOL.hidemem.Value ? "collapse-hidden" : ""
            },
            React.createElement(
              "table",
              { className: "table table-striped"
              },
              React.createElement(
                "thead",
                null,
                React.createElement(
                  "tr",
                  null,
                  React.createElement("th", null),
                  React.createElement(
                    "th",
                    { className: "text-right"
                    },
                    "Free"
                  ),
                  React.createElement(
                    "th",
                    { className: "text-right"
                    },
                    "Used"
                  ),
                  React.createElement(
                    "th",
                    { className: "text-right"
                    },
                    "Total"
                  )
                )
              ),
              React.createElement(
                "tbody",
                null,
                rows
              )
            )
          )
        )
      );
    },

    ifbytes_rows: function ifbytes_rows(Data, $if) {
      return React.createElement(
        "tr",
        { key: "ifbytes-rowby-name-" + $if.Name
        },
        React.createElement(
          "td",
          null,
          $if.Name
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.DeltaIn
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.DeltaOut
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.In
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.Out
        )
      );
    },
    iferrors_rows: function iferrors_rows(Data, $if) {
      return React.createElement(
        "tr",
        { key: "iferrors-rowby-name-" + $if.Name
        },
        React.createElement(
          "td",
          null,
          $if.Name
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.DeltaIn
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.DeltaOut
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.In
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.Out
        )
      );
    },
    ifpackets_rows: function ifpackets_rows(Data, $if) {
      return React.createElement(
        "tr",
        { key: "ifpackets-rowby-name-" + $if.Name
        },
        React.createElement(
          "td",
          null,
          $if.Name
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.DeltaIn
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.DeltaOut
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.In
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $if.Out
        )
      );
    },
    panelif: function panelif(Data, r1, r2, r3) {
      return React.createElement(
        "div",
        null,
        React.createElement(
          "div",
          null,
          React.createElement(
            "a",
            { href: Data.Params.BOOL.configif.Href, onClick: this.handleClick, className: "btn-block"
            },
            React.createElement(
              "span",
              { className: Data.Params.BOOL.configif.Value ? "h4 bg-info" : "h4"
              },
              Data.Params.ENUM.ift.Title
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { id: "ifconfig", className: Data.Params.BOOL.configif.Value ? "config-margintop" : "config-margintop collapse-hidden"
            },
            React.createElement(
              "form",
              { className: "form-inline", action: "/form/" + Data.Params.Query
              },
              React.createElement("input", { className: "hidden-submit", type: "submit"
              }),
              React.createElement(
                "div",
                { className: "btn-toolbar"
                },
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "a",
                    { className: Data.Params.BOOL.hideif.Value ? "btn btn-default active" : "btn btn-default", href: Data.Params.BOOL.hideif.Href, onClick: this.handleClick
                    },
                    "Hidden"
                  ),
                  React.createElement(
                    "a",
                    { className: Data.ExpandableIF ? "btn btn-default" : "btn btn-default disabled", href: Data.Params.BOOL.expandif.Href, onClick: this.handleClick
                    },
                    Data.ExpandtextIF
                  )
                ),
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "div",
                    { className: "input-group input-group-sm refresh-group" + (Data.Params.PERIOD.refreshif.InputErrd ? " has-warning" : "")
                    },
                    React.createElement(
                      "span",
                      { className: "input-group-addon"
                      },
                      "Refresh"
                    ),
                    React.createElement("input", { className: "form-control refresh-input width-fourem", type: "text", placeholder: Data.Params.PERIOD.refreshif.Placeholder, name: "refreshif", value: Data.Params.PERIOD.refreshif.Input, onChange: this.handleChange
                    })
                  )
                )
              )
            ),
            React.createElement(
              "ul",
              { className: "nav nav-tabs config-margintop"
              },
              React.createElement(
                "li",
                { className: Data.Params.ENUM.ift.Uint == 0 ? "active" : ""
                },
                React.createElement(
                  "a",
                  { href: Data.Params.ENUM.ift.PACKETS.Href, onClick: this.handleClick
                  },
                  "Packets"
                )
              ),
              React.createElement(
                "li",
                { className: Data.Params.ENUM.ift.Uint == 1 ? "active" : ""
                },
                React.createElement(
                  "a",
                  { href: Data.Params.ENUM.ift.ERRORS.Href, onClick: this.handleClick
                  },
                  "Errors"
                )
              ),
              React.createElement(
                "li",
                { className: Data.Params.ENUM.ift.Uint == 2 ? "active" : ""
                },
                React.createElement(
                  "a",
                  { href: Data.Params.ENUM.ift.IFBYTES.Href, onClick: this.handleClick
                  },
                  "Bytes"
                )
              )
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { className: Data.Params.BOOL.hideif.Value ? "collapse-hidden" : ""
            },
            React.createElement(
              "div",
              { className: Data.Params.ENUM.ift.Uint == 0 ? "" : "collapse-hidden"
              },
              React.createElement(
                "table",
                { className: "table table-striped"
                },
                React.createElement(
                  "thead",
                  null,
                  React.createElement(
                    "tr",
                    null,
                    React.createElement(
                      "th",
                      null,
                      "Interface"
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "per second"
                      },
                      "In ",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        "ps"
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "per second"
                      },
                      "Out ",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        "ps"
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "total modulo 4G"
                      },
                      "In ",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        "%4G"
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "total modulo 4G"
                      },
                      "Out ",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        "%4G"
                      )
                    )
                  )
                ),
                React.createElement(
                  "tbody",
                  null,
                  r1
                )
              )
            ),
            React.createElement(
              "div",
              { className: Data.Params.ENUM.ift.Uint == 1 ? "" : "collapse-hidden"
              },
              React.createElement(
                "table",
                { className: "table table-striped"
                },
                React.createElement(
                  "thead",
                  null,
                  React.createElement(
                    "tr",
                    null,
                    React.createElement(
                      "th",
                      null,
                      "Interface"
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "per second"
                      },
                      "In ",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        "ps"
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "per second"
                      },
                      "Out ",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        "ps"
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "modulo 4G"
                      },
                      "In ",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        "%4G"
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "modulo 4G"
                      },
                      "Out ",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        "%4G"
                      )
                    )
                  )
                ),
                React.createElement(
                  "tbody",
                  null,
                  r2
                )
              )
            ),
            React.createElement(
              "div",
              { className: Data.Params.ENUM.ift.Uint == 2 ? "" : "collapse-hidden"
              },
              React.createElement(
                "table",
                { className: "table table-striped"
                },
                React.createElement(
                  "thead",
                  null,
                  React.createElement(
                    "tr",
                    null,
                    React.createElement(
                      "th",
                      null,
                      "Interface"
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "BITS per second"
                      },
                      "In",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        React.createElement(
                          "i",
                          null,
                          "b"
                        ),
                        "ps"
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "BITS per second"
                      },
                      "Out",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        React.createElement(
                          "i",
                          null,
                          "b"
                        ),
                        "ps"
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "total BYTES modulo 4G"
                      },
                      "In",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        React.createElement(
                          "i",
                          null,
                          "B"
                        ),
                        "%4G"
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "text-right nowrap", title: "total BYTES modulo 4G"
                      },
                      "Out",
                      React.createElement(
                        "span",
                        { className: "unit"
                        },
                        React.createElement(
                          "i",
                          null,
                          "B"
                        ),
                        "%4G"
                      )
                    )
                  )
                ),
                React.createElement(
                  "tbody",
                  null,
                  r3
                )
              )
            )
          )
        )
      );
    },

    cpu_rows: function cpu_rows(Data, $core) {
      return React.createElement(
        "tr",
        { key: "cpu-rowby-N-" + $core.N
        },
        React.createElement(
          "td",
          { className: "text-right nowrap"
          },
          $core.N
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          React.createElement(
            "span",
            { className: "usepercent-text", "data-usepercent": $core.User
            },
            $core.User
          )
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          React.createElement(
            "span",
            { className: "usepercent-text", "data-usepercent": $core.Sys
            },
            $core.Sys
          )
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          React.createElement(
            "span",
            { className: "usepercent-text-inverse", "data-usepercent": $core.Idle
            },
            $core.Idle
          )
        )
      );
    },
    panelcpu: function panelcpu(Data, rows) {
      return React.createElement(
        "div",
        null,
        React.createElement(
          "div",
          null,
          React.createElement(
            "a",
            { href: Data.Params.BOOL.configcpu.Href, onClick: this.handleClick, className: "btn-block"
            },
            React.createElement(
              "span",
              { className: Data.Params.BOOL.configcpu.Value ? "h4 bg-info" : "h4"
              },
              "CPU"
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { id: "cpuconfig", className: Data.Params.BOOL.configcpu.Value ? "config-margintop" : "config-margintop collapse-hidden"
            },
            React.createElement(
              "form",
              { className: "form-inline", action: "/form/" + Data.Params.Query
              },
              React.createElement("input", { className: "hidden-submit", type: "submit"
              }),
              React.createElement(
                "div",
                { className: "btn-toolbar"
                },
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "a",
                    { className: Data.Params.BOOL.hidecpu.Value ? "btn btn-default active" : "btn btn-default", href: Data.Params.BOOL.hidecpu.Href, onClick: this.handleClick
                    },
                    "Hidden"
                  ),
                  React.createElement(
                    "a",
                    { className: Data.CPU.ExpandableCPU ? "btn btn-default" : "btn btn-default disabled", href: Data.Params.BOOL.expandcpu.Href, onClick: this.handleClick
                    },
                    Data.CPU.ExpandtextCPU
                  )
                ),
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "div",
                    { className: "input-group input-group-sm refresh-group" + (Data.Params.PERIOD.refreshcpu.InputErrd ? " has-warning" : "")
                    },
                    React.createElement(
                      "span",
                      { className: "input-group-addon"
                      },
                      "Refresh"
                    ),
                    React.createElement("input", { className: "form-control refresh-input width-fourem", type: "text", placeholder: Data.Params.PERIOD.refreshcpu.Placeholder, name: "refreshcpu", value: Data.Params.PERIOD.refreshcpu.Input, onChange: this.handleChange
                    })
                  )
                )
              )
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { className: Data.Params.BOOL.hidecpu.Value ? "collapse-hidden" : ""
            },
            React.createElement(
              "table",
              { className: "table table-striped"
              },
              React.createElement(
                "thead",
                null,
                React.createElement(
                  "tr",
                  null,
                  React.createElement("th", null),
                  React.createElement(
                    "th",
                    { className: "text-right nowrap"
                    },
                    "User",
                    React.createElement(
                      "span",
                      { className: "unit"
                      },
                      "%"
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "text-right nowrap"
                    },
                    "Sys",
                    React.createElement(
                      "span",
                      { className: "unit"
                      },
                      "%"
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "text-right nowrap"
                    },
                    "Idle",
                    React.createElement(
                      "span",
                      { className: "unit"
                      },
                      "%"
                    )
                  )
                )
              ),
              React.createElement(
                "tbody",
                null,
                rows
              )
            )
          )
        )
      );
    },

    dfbytes_rows: function dfbytes_rows(Data, $disk) {
      return React.createElement(
        "tr",
        { key: "dfbytes-rowby-dirname-" + $disk.DirName
        },
        React.createElement(
          "td",
          null,
          $disk.DevName
        ),
        React.createElement(
          "td",
          null,
          $disk.DirName
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $disk.Avail
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          React.createElement(
            "span",
            { className: "label", "data-usepercent": $disk.UsePercent
            },
            $disk.UsePercent,
            "%"
          ),
          " ",
          $disk.Used
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $disk.Total
        )
      );
    },
    dfinodes_rows: function dfinodes_rows(Data, $disk) {
      return React.createElement(
        "tr",
        { key: "dfinodes-rowby-dirname-" + $disk.DirName
        },
        React.createElement(
          "td",
          null,
          $disk.DevName
        ),
        React.createElement(
          "td",
          null,
          $disk.DirName
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $disk.Ifree
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          React.createElement(
            "span",
            { className: "label", "data-usepercent": $disk.IusePercent
            },
            $disk.IusePercent,
            "%"
          ),
          " ",
          $disk.Iused
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          $disk.Inodes
        )
      );
    },
    paneldf: function paneldf(Data, r1, r2) {
      return React.createElement(
        "div",
        null,
        React.createElement(
          "div",
          null,
          React.createElement(
            "a",
            { href: Data.Params.BOOL.configdf.Href, onClick: this.handleClick, className: "btn-block"
            },
            React.createElement(
              "span",
              { className: Data.Params.BOOL.configdf.Value ? "h4 bg-info" : "h4"
              },
              Data.Params.ENUM.dft.Title
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { id: "dfconfig", className: Data.Params.BOOL.configdf.Value ? "config-margintop" : "config-margintop collapse-hidden"
            },
            React.createElement(
              "form",
              { className: "form-inline", action: "/form/" + Data.Params.Query
              },
              React.createElement("input", { className: "hidden-submit", type: "submit"
              }),
              React.createElement(
                "div",
                { className: "btn-toolbar"
                },
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "a",
                    { className: Data.Params.BOOL.hidedf.Value ? "btn btn-default active" : "btn btn-default", href: Data.Params.BOOL.hidedf.Href, onClick: this.handleClick
                    },
                    "Hidden"
                  ),
                  React.createElement(
                    "a",
                    { className: Data.ExpandableDF ? "btn btn-default" : "btn btn-default disabled", href: Data.Params.BOOL.expanddf.Href, onClick: this.handleClick
                    },
                    Data.ExpandtextDF
                  )
                ),
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "div",
                    { className: "input-group input-group-sm refresh-group" + (Data.Params.PERIOD.refreshdf.InputErrd ? " has-warning" : "")
                    },
                    React.createElement(
                      "span",
                      { className: "input-group-addon"
                      },
                      "Refresh"
                    ),
                    React.createElement("input", { className: "form-control refresh-input width-fourem", type: "text", placeholder: Data.Params.PERIOD.refreshdf.Placeholder, name: "refreshdf", value: Data.Params.PERIOD.refreshdf.Input, onChange: this.handleChange
                    })
                  )
                )
              )
            ),
            React.createElement(
              "ul",
              { className: "nav nav-tabs config-margintop"
              },
              React.createElement(
                "li",
                { className: Data.Params.ENUM.dft.Uint == 0 ? "active" : ""
                },
                React.createElement(
                  "a",
                  { href: Data.Params.ENUM.dft.INODES.Href, onClick: this.handleClick
                  },
                  "Inodes"
                )
              ),
              React.createElement(
                "li",
                { className: Data.Params.ENUM.dft.Uint == 1 ? "active" : ""
                },
                React.createElement(
                  "a",
                  { href: Data.Params.ENUM.dft.DFBYTES.Href, onClick: this.handleClick
                  },
                  "Bytes"
                )
              )
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { className: Data.Params.BOOL.hidedf.Value ? "collapse-hidden" : ""
            },
            React.createElement(
              "div",
              { className: Data.Params.ENUM.dft.Uint == 0 ? "" : "collapse-hidden"
              },
              React.createElement(
                "table",
                { className: "table table-striped"
                },
                React.createElement(
                  "thead",
                  null,
                  React.createElement(
                    "tr",
                    null,
                    React.createElement(
                      "th",
                      { className: "header"
                      },
                      "Device"
                    ),
                    React.createElement(
                      "th",
                      { className: "header"
                      },
                      "Mounted"
                    ),
                    React.createElement(
                      "th",
                      { className: "header text-right"
                      },
                      "Avail"
                    ),
                    React.createElement(
                      "th",
                      { className: "header text-right"
                      },
                      "Used"
                    ),
                    React.createElement(
                      "th",
                      { className: "header text-right"
                      },
                      "Total"
                    )
                  )
                ),
                React.createElement(
                  "tbody",
                  null,
                  r1
                )
              )
            ),
            React.createElement(
              "div",
              { className: Data.Params.ENUM.dft.Uint == 1 ? "" : "collapse-hidden"
              },
              React.createElement(
                "table",
                { className: "table table-striped"
                },
                React.createElement(
                  "thead",
                  null,
                  React.createElement(
                    "tr",
                    null,
                    React.createElement(
                      "th",
                      { className: "header "
                      },
                      React.createElement(
                        "a",
                        { href: Data.Params.ENUM.df.FS.Href, className: Data.Params.ENUM.df.FS.Class
                        },
                        "Device",
                        React.createElement("span", { className: Data.Params.ENUM.df.FS.CaretClass
                        })
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "header "
                      },
                      React.createElement(
                        "a",
                        { href: Data.Params.ENUM.df.MP.Href, className: Data.Params.ENUM.df.MP.Class
                        },
                        "Mounted",
                        React.createElement("span", { className: Data.Params.ENUM.df.MP.CaretClass
                        })
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "header text-right"
                      },
                      React.createElement(
                        "a",
                        { href: Data.Params.ENUM.df.AVAIL.Href, className: Data.Params.ENUM.df.AVAIL.Class
                        },
                        "Avail",
                        React.createElement("span", { className: Data.Params.ENUM.df.AVAIL.CaretClass
                        })
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "header text-right"
                      },
                      React.createElement(
                        "a",
                        { href: Data.Params.ENUM.df.USED.Href, className: Data.Params.ENUM.df.USED.Class
                        },
                        "Used",
                        React.createElement("span", { className: Data.Params.ENUM.df.USED.CaretClass
                        })
                      )
                    ),
                    React.createElement(
                      "th",
                      { className: "header text-right"
                      },
                      React.createElement(
                        "a",
                        { href: Data.Params.ENUM.df.TOTAL.Href, className: Data.Params.ENUM.df.TOTAL.Class
                        },
                        "Total",
                        React.createElement("span", { className: Data.Params.ENUM.df.TOTAL.CaretClass
                        })
                      )
                    )
                  )
                ),
                React.createElement(
                  "tbody",
                  null,
                  r2
                )
              )
            )
          )
        )
      );
    },

    ps_rows: function ps_rows(Data, $proc) {
      return React.createElement(
        "tr",
        { key: "ps-rowby-pid-" + $proc.PID
        },
        React.createElement(
          "td",
          { className: "text-right"
          },
          " ",
          $proc.PID
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          " ",
          $proc.UID
        ),
        React.createElement(
          "td",
          null,
          $proc.User
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          " ",
          $proc.Priority
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          " ",
          $proc.Nice
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          " ",
          $proc.Size
        ),
        React.createElement(
          "td",
          { className: "text-right"
          },
          " ",
          $proc.Resident
        ),
        React.createElement(
          "td",
          { className: "text-center"
          },
          $proc.Time
        ),
        React.createElement(
          "td",
          null,
          $proc.Name
        )
      );
    },
    panelps: function panelps(Data, rows) {
      return React.createElement(
        "div",
        null,
        React.createElement(
          "div",
          null,
          React.createElement(
            "a",
            { href: Data.Params.BOOL.configps.Href, onClick: this.handleClick, className: "btn-block"
            },
            React.createElement(
              "span",
              { className: Data.Params.BOOL.configps.Value ? "h4 bg-info" : "h4"
              },
              "Processes"
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { id: "psconfig", className: Data.Params.BOOL.configps.Value ? "config-margintop" : "config-margintop collapse-hidden"
            },
            React.createElement(
              "form",
              { className: "form-inline text-right", action: "/form/" + Data.Params.Query
              },
              React.createElement("input", { className: "hidden-submit", type: "submit"
              }),
              React.createElement(
                "div",
                { className: "btn-toolbar"
                },
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "a",
                    { className: Data.Params.BOOL.hideps.Value ? "btn btn-default active" : "btn btn-default", href: Data.Params.BOOL.hideps.Href, onClick: this.handleClick
                    },
                    "Hidden"
                  ),
                  React.createElement(
                    "a",
                    { className: Data.PStable.PSnotDecreasable ? "btn btn-default" : "btn btn-default disabled", href: Data.Params.LIMIT.psn.LessHref, onClick: this.handleClick
                    },
                    "-"
                  ),
                  React.createElement(
                    "a",
                    { className: Data.PStable.PSnotExpandable ? "btn btn-default" : "btn btn-default disabled", href: Data.Params.LIMIT.psn.MoreHref, onClick: this.handleClick
                    },
                    Data.PStable.PSplusText
                  )
                ),
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "div",
                    { className: "input-group input-group-sm refresh-group" + (Data.Params.PERIOD.refreshps.InputErrd ? " has-warning" : "")
                    },
                    React.createElement(
                      "span",
                      { className: "input-group-addon"
                      },
                      "Refresh"
                    ),
                    React.createElement("input", { className: "form-control refresh-input width-fourem", type: "text", placeholder: Data.Params.PERIOD.refreshps.Placeholder, name: "refreshps", value: Data.Params.PERIOD.refreshps.Input, onChange: this.handleChange
                    })
                  )
                )
              )
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { className: Data.Params.BOOL.hideps.Value ? "collapse-hidden" : ""
            },
            React.createElement(
              "table",
              { className: "table table-striped table-hover"
              },
              React.createElement(
                "thead",
                null,
                React.createElement(
                  "tr",
                  null,
                  React.createElement(
                    "th",
                    { className: "header text-right"
                    },
                    React.createElement(
                      "a",
                      { href: Data.Params.ENUM.ps.PID.Href, className: Data.Params.ENUM.ps.PID.Class
                      },
                      "PID",
                      React.createElement("span", { className: Data.Params.ENUM.ps.PID.CaretClass
                      })
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "header text-right"
                    },
                    React.createElement(
                      "a",
                      { href: Data.Params.ENUM.ps.UID.Href, className: Data.Params.ENUM.ps.UID.Class
                      },
                      "UID",
                      React.createElement("span", { className: Data.Params.ENUM.ps.UID.CaretClass
                      })
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "header "
                    },
                    React.createElement(
                      "a",
                      { href: Data.Params.ENUM.ps.USER.Href, className: Data.Params.ENUM.ps.USER.Class
                      },
                      "USER",
                      React.createElement("span", { className: Data.Params.ENUM.ps.USER.CaretClass
                      })
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "header text-right"
                    },
                    React.createElement(
                      "a",
                      { href: Data.Params.ENUM.ps.PRI.Href, className: Data.Params.ENUM.ps.PRI.Class
                      },
                      "PR",
                      React.createElement("span", { className: Data.Params.ENUM.ps.PRI.CaretClass
                      })
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "header text-right"
                    },
                    React.createElement(
                      "a",
                      { href: Data.Params.ENUM.ps.NICE.Href, className: Data.Params.ENUM.ps.NICE.Class
                      },
                      "NI",
                      React.createElement("span", { className: Data.Params.ENUM.ps.NICE.CaretClass
                      })
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "header text-right"
                    },
                    React.createElement(
                      "a",
                      { href: Data.Params.ENUM.ps.VIRT.Href, className: Data.Params.ENUM.ps.VIRT.Class
                      },
                      "VIRT",
                      React.createElement("span", { className: Data.Params.ENUM.ps.VIRT.CaretClass
                      })
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "header text-right"
                    },
                    React.createElement(
                      "a",
                      { href: Data.Params.ENUM.ps.RES.Href, className: Data.Params.ENUM.ps.RES.Class
                      },
                      "RES",
                      React.createElement("span", { className: Data.Params.ENUM.ps.RES.CaretClass
                      })
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "header text-center"
                    },
                    React.createElement(
                      "a",
                      { href: Data.Params.ENUM.ps.TIME.Href, className: Data.Params.ENUM.ps.TIME.Class
                      },
                      "TIME",
                      React.createElement("span", { className: Data.Params.ENUM.ps.TIME.CaretClass
                      })
                    )
                  ),
                  React.createElement(
                    "th",
                    { className: "header "
                    },
                    React.createElement(
                      "a",
                      { href: Data.Params.ENUM.ps.NAME.Href, className: Data.Params.ENUM.ps.NAME.Class
                      },
                      "COMMAND",
                      React.createElement("span", { className: Data.Params.ENUM.ps.NAME.CaretClass
                      })
                    )
                  )
                )
              ),
              React.createElement(
                "tbody",
                null,
                rows
              )
            )
          )
        )
      );
    },

    vagrant_rows: function vagrant_rows(Data, $mach) {
      return React.createElement(
        "tr",
        { key: "vagrant-rowby-uuid-" + $mach.UUID
        },
        React.createElement(
          "td",
          null,
          $mach.UUID
        ),
        React.createElement(
          "td",
          null,
          $mach.Name
        ),
        React.createElement(
          "td",
          null,
          $mach.Provider
        ),
        React.createElement(
          "td",
          null,
          $mach.State
        ),
        React.createElement(
          "td",
          null,
          $mach.Vagrantfile_path
        )
      );
    },
    vagrant_error: function vagrant_error(Data) {
      return React.createElement(
        "tr",
        { key: "vgerror"
        },
        React.createElement(
          "td",
          { colSpan: "5"
          },
          Data.VagrantError
        )
      );
    },
    panelvg: function panelvg(Data, rows) {
      return React.createElement(
        "div",
        null,
        React.createElement(
          "div",
          null,
          React.createElement(
            "a",
            { href: Data.Params.BOOL.configvg.Href, onClick: this.handleClick, className: "btn-block"
            },
            React.createElement(
              "span",
              { className: Data.Params.BOOL.configvg.Value ? "h4 bg-info" : "h4"
              },
              "Vagrant global-status"
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { id: "vgconfig", className: Data.Params.BOOL.configvg.Value ? "config-margintop" : "config-margintop collapse-hidden"
            },
            React.createElement(
              "form",
              { className: "form-inline", action: "/form/" + Data.Params.Query
              },
              React.createElement("input", { className: "hidden-submit", type: "submit"
              }),
              React.createElement(
                "div",
                { className: "btn-toolbar"
                },
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "a",
                    { className: Data.Params.BOOL.hidevg.Value ? "btn btn-default active" : "btn btn-default", href: Data.Params.BOOL.hidevg.Href, onClick: this.handleClick
                    },
                    "Hidden"
                  )
                ),
                React.createElement(
                  "div",
                  { className: "btn-group btn-group-sm", role: "group"
                  },
                  React.createElement(
                    "div",
                    { className: "input-group input-group-sm refresh-group" + (Data.Params.PERIOD.refreshvg.InputErrd ? " has-warning" : "")
                    },
                    React.createElement(
                      "span",
                      { className: "input-group-addon"
                      },
                      "Refresh"
                    ),
                    React.createElement("input", { className: "form-control refresh-input width-fourem", type: "text", placeholder: Data.Params.PERIOD.refreshvg.Placeholder, name: "refreshvg", value: Data.Params.PERIOD.refreshvg.Input, onChange: this.handleChange
                    })
                  )
                )
              )
            )
          )
        ),
        React.createElement(
          "div",
          null,
          React.createElement(
            "div",
            { className: Data.Params.BOOL.hidevg.Value ? "collapse-hidden" : ""
            },
            React.createElement(
              "table",
              { id: "vgtable", className: "table table-striped"
              },
              React.createElement(
                "thead",
                null,
                React.createElement(
                  "tr",
                  null,
                  React.createElement(
                    "th",
                    null,
                    "id"
                  ),
                  React.createElement(
                    "th",
                    null,
                    "name"
                  ),
                  React.createElement(
                    "th",
                    null,
                    "provider"
                  ),
                  React.createElement(
                    "th",
                    null,
                    "state"
                  ),
                  React.createElement(
                    "th",
                    null,
                    "directory"
                  )
                )
              ),
              React.createElement(
                "tbody",
                null,
                rows
              )
            )
          )
        )
      );
    }
  };
});
