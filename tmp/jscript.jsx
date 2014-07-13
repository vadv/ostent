/** @jsx React.DOM */ function mem_rows       (Data, $mem)  { return (<tr key={$mem.Kind}><td>{$mem.Kind}</td><td className="text-right">{$mem.Free}</td><td className="text-right">{$mem.Used}&nbsp;<sup><span dangerouslySetInnerHTML={{__html: $mem.UsePercentHTML}} /></sup></td><td className="text-right">{$mem.Total}</td></tr>); }
function mem_table      (Data, rows)  { return (<table className="table1 stripe-table"><thead><tr><th></th><th className="text-right">Free</th><th className="text-right">Used</th><th className="text-right">Total</th></tr></thead><tbody>{rows}</tbody></table>); }

function ifbytes_rows   (Data, $if)   { return (<tr key={$if.NameKey}><td><span dangerouslySetInnerHTML={{__html: $if.NameHTML}} /></td><td className="text-right">{$if.DeltaIn}</td><td className="text-right">{$if.DeltaOut}</td><td className="text-right">{$if.In}</td><td className="text-right">{$if.Out}</td></tr>); }
function ifbytes_table  (Data, rows)  { return (<table className="table1 stripe-table"><thead><tr><th>Interface</th><th className="text-right nowrap" title="BITS per second">In<span className="unit"><i>b</i>ps</span></th><th className="text-right nowrap" title="BITS per second">Out<span className="unit"><i>b</i>ps</span></th><th className="text-right nowrap" title="total BYTES modulo 4G">In<span className="unit"><i>B</i>%4G</span></th><th className="text-right nowrap" title="total BYTES modulo 4G">Out<span className="unit"><i>B</i>%4G</span></th></tr></thead><tbody>{rows}</tbody></table>); }
function iferrors_rows  (Data, $if)   { return (<tr key={$if.NameKey}><td><span dangerouslySetInnerHTML={{__html: $if.NameHTML}} /></td><td className="text-right">{$if.DeltaIn}</td><td className="text-right">{$if.DeltaOut}</td><td className="text-right">{$if.In}</td><td className="text-right">{$if.Out}</td></tr>); }
function iferrors_table (Data, rows)  { return (<table className="table1 stripe-table"><thead><tr><th>Interface</th><th className="text-right nowrap" title="per second">In&nbsp;<span className="unit">ps</span></th><th className="text-right nowrap" title="per second">Out&nbsp;<span className="unit">ps</span></th><th className="text-right nowrap" title="modulo 4G">In&nbsp;<span className="unit">%4G</span></th><th className="text-right nowrap" title="modulo 4G">Out&nbsp;<span className="unit">%4G</span></th></tr></thead><tbody>{rows}</tbody></table>); }
function ifpackets_rows (Data, $if)   { return (<tr key={$if.NameKey}><td><span dangerouslySetInnerHTML={{__html: $if.NameHTML}} /></td><td className="text-right">{$if.DeltaIn}</td><td className="text-right">{$if.DeltaOut}</td><td className="text-right">{$if.In}</td><td className="text-right">{$if.Out}</td></tr>); }
function ifpackets_table(Data, rows)  { return (<table className="table1 stripe-table"><thead><tr><th>Interface</th><th className="text-right nowrap" title="per second">In&nbsp;<span className="unit">ps</span></th><th className="text-right nowrap" title="per second">Out&nbsp;<span className="unit">ps</span></th><th title="total modulo 4G" className="text-right nowrap">In&nbsp;<span className="unit">%4G</span></th><th className="text-right nowrap" title="total modulo 4G">Out&nbsp;<span className="unit">%4G</span></th></tr></thead><tbody>{rows}</tbody></table>); }

function cpu_rows       (Data, $core) { return (<tr key={$core.N}><td className="text-right nowrap">{$core.N}</td><td className="text-right"><span className={$core.UserClass}>{$core.User}</span></td><td className="text-right"><span className={$core.SysClass}>{$core.Sys}</span></td><td className="text-right"><span className={$core.IdleClass}>{$core.Idle}</span></td></tr>); }
function cpu_table      (Data, rows)  { return (<table className="table1 stripe-table"><thead><tr><th></th><th className="text-right nowrap">User<span className="unit">%</span></th><th className="text-right nowrap">Sys<span className="unit">%</span></th><th className="text-right nowrap">Idle<span className="unit">%</span></th></tr></thead><tbody>{rows}</tbody></table>); }

function dfbytes_rows   (Data, $disk) { return (<tr key={$disk.DirNameKey}><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $disk.DiskNameHTML}} /></td><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $disk.DirNameHTML}} /></td><td className="text-right">{$disk.Avail}</td><td className="text-right">{$disk.Used}&nbsp;<sup><span className={$disk.UsePercentClass}>{$disk.UsePercent}%</span></sup></td><td className="text-right">{$disk.Total}</td></tr>); }
function dfbytes_table  (Data, rows)  { return (<table className="table1 stripe-table"><thead><tr><th className="header">        <a href={Data.DFlinks.DiskName.Href} className={Data.DFlinks.DiskName.Class}>Device<span  className={Data.DFlinks.DiskName.CaretClass} /></a></th><th className="header">        <a href={Data.DFlinks.DirName.Href}  className={Data.DFlinks.DirName.Class} >Mounted<span className={Data.DFlinks.DirName.CaretClass}  /></a></th><th className="header text-right"><a href={Data.DFlinks.Avail.Href}    className={Data.DFlinks.Avail.Class}   >Avail<span   className={Data.DFlinks.Avail.CaretClass}    /></a></th><th className="header text-right"><a href={Data.DFlinks.Used.Href}     className={Data.DFlinks.Used.Class}    >Used<span    className={Data.DFlinks.Used.CaretClass}     /></a></th><th className="header text-right"><a href={Data.DFlinks.Total.Href}    className={Data.DFlinks.Total.Class}   >Total<span   className={Data.DFlinks.Total.CaretClass}    /></a></th></tr></thead><tbody>{rows}</tbody></table>); }
function dfinodes_rows  (Data, $disk) { return (<tr key={$disk.DirNameKey}><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $disk.DiskNameHTML}} /></td><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $disk.DirNameHTML}} /></td><td className="text-right">{$disk.Ifree}</td><td className="text-right">{$disk.Iused}&nbsp;<sup><span className={$disk.IusePercentClass}>{$disk.IusePercent}%</span></sup></td><td className="text-right">{$disk.Inodes}</td></tr>); }
function dfinodes_table (Data, rows)  { return (<table className="table1 stripe-table"><thead><tr><th className="header">Device</th><th className="header">Mounted</th><th className="header text-right">Avail</th><th className="header text-right">Used</th><th className="header text-right">Total</th></tr></thead><tbody>{rows}</tbody></table>); }

function ps_rows        (Data, $proc) { return (<tr key={$proc.PID}><td className="text-right">{$proc.PID}</td><td className="text-right"><span dangerouslySetInnerHTML={{__html: $proc.UserHTML}} /></td><td className="text-right">{$proc.Priority}</td><td className="text-right">{$proc.Nice}</td><td className="text-right">{$proc.Size}</td><td className="text-right">{$proc.Resident}</td><td className="text-center">{$proc.Time}</td><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $proc.NameHTML}} /></td></tr>); }
function ps_table       (Data, rows)  { return (<table className="table2 stripe-table"><thead><tr><th className="header text-right"> <a href={Data.PSlinks.PID.Href}      className={Data.PSlinks.PID.Class}     >PID<span     className={Data.PSlinks.PID.CaretClass}      /></a></th><th className="header text-right"> <a href={Data.PSlinks.User.Href}     className={Data.PSlinks.User.Class}    >USER<span    className={Data.PSlinks.User.CaretClass}     /></a></th><th className="header text-right"> <a href={Data.PSlinks.Priority.Href} className={Data.PSlinks.Priority.Class}>PR<span      className={Data.PSlinks.Priority.CaretClass} /></a></th><th className="header text-right"> <a href={Data.PSlinks.Nice.Href}     className={Data.PSlinks.Nice.Class}    >NI<span      className={Data.PSlinks.Nice.CaretClass}     /></a></th><th className="header text-right"> <a href={Data.PSlinks.Size.Href}     className={Data.PSlinks.Size.Class}    >VIRT<span    className={Data.PSlinks.Size.CaretClass}     /></a></th><th className="header text-right"> <a href={Data.PSlinks.Resident.Href} className={Data.PSlinks.Resident.Class}>RES<span     className={Data.PSlinks.Resident.CaretClass} /></a></th><th className="header text-center"><a href={Data.PSlinks.Time.Href}     className={Data.PSlinks.Time.Class}    >TIME<span    className={Data.PSlinks.Time.CaretClass}     /></a></th><th className="header">            <a href={Data.PSlinks.Name.Href}     className={Data.PSlinks.Name.Class}    >COMMAND<span className={Data.PSlinks.Name.CaretClass}     /></a></th></tr></thead><tbody>{rows}</tbody></table>); }

function vagrant_rows(Data, $machine) { return (<tr key={$machine.UUID}><td><span dangerouslySetInnerHTML={{__html: $machine.UUIDHTML}} /></td><td>{$machine.Name}</td><td>{$machine.Provider}</td><td>{$machine.State}</td><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $machine.Vagrantfile_pathHTML}} /></td></tr>); }
function vagrant_error  (Data)        { return (<tr key="vagrant-error"><td colspan="5">{Data.VagrantError}</td></tr>); }
function vagrant_table  (Data, rows)  { return (<table className="table1 stripe-table" id="vagrant-table"><thead><tr><th>id</th><th>name</th><th>provider</th><th>state</th><th>directory</th></tr></thead><tbody>{rows}</tbody></table>); }
