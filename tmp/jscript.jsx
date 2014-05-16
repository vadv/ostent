/** @jsx React.DOM */ function emptyK(obj, key) {
return (obj      === undefined ||
	obj      === null      ||
	obj[key] === undefined ||
	obj[key] === null);
}

var IFbytesCLASS = React.createClass({
	getInitialState: function() { return Data.IFbytes; },

	render: function() {
		var Data = {IFbytes: this.state};
		var rows = emptyK(Data.IFbytes, 'List') ?'': Data.IFbytes.List.map(function($if) {
			return (<tr key={$if.NameKey}><td><span dangerouslySetInnerHTML={{__html: $if.NameHTML}} /></td><td className="text-right">{$if.DeltaIn}</td><td className="text-right">{$if.DeltaOut}</td><td className="text-right">{$if.In}</td><td className="text-right">{$if.Out}</td></tr>);
		});
		
		return (<table className="table1 stripe-table"><thead><tr><th>Interface</th><th className="text-right nowrap" title="BITS per second">In<span className="unit"><b>b</b>ps</span></th><th className="text-right nowrap" title="BITS per second">Out<span className="unit"><b>b</b>ps</span></th><th className="text-right nowrap" title="total BYTES modulo 4G">In<span className="unit"><b>B</b>%4G</span></th><th className="text-right nowrap" title="total BYTES modulo 4G">Out<span className="unit"><b>B</b>%4G</span></th></tr></thead><tbody>{rows}</tbody></table>);
		
	}
});

var IFerrorsCLASS = React.createClass({
	getInitialState: function() { return Data.IFerrors; },

	render: function() {
		var Data = {IFerrors: this.state};
		var rows = emptyK(Data.IFerrors, 'List') ?'': Data.IFerrors.List.map(function($if) {
			return (<tr key={$if.NameKey}><td><span dangerouslySetInnerHTML={{__html: $if.NameHTML}} /></td><td className="text-right">{$if.DeltaIn}</td><td className="text-right">{$if.DeltaOut}</td><td className="text-right">{$if.In}</td><td className="text-right">{$if.Out}</td></tr>);
		});
		
		return (<table className="table1 stripe-table"><thead><tr><th>Interface</th><th className="text-right nowrap" title="per second">In&nbsp;<span className="unit">ps</span></th><th className="text-right nowrap" title="per second">Out&nbsp;<span className="unit">ps</span></th><th className="text-right nowrap" title="modulo 4G">In&nbsp;<span className="unit">%4G</span></th><th className="text-right nowrap" title="modulo 4G">Out&nbsp;<span className="unit">%4G</span></th></tr></thead><tbody>{rows}</tbody></table>);
		
	}
});

var IFpacketsCLASS = React.createClass({
	getInitialState: function() { return Data.IFpackets; },

	render: function() {
		var Data = {IFpackets: this.state};
		var rows = emptyK(Data.IFpackets, 'List') ?'': Data.IFpackets.List.map(function($if) {
			return (<tr key={$if.NameKey}><td><span dangerouslySetInnerHTML={{__html: $if.NameHTML}} /></td><td className="text-right">{$if.DeltaIn}</td><td className="text-right">{$if.DeltaOut}</td><td className="text-right">{$if.In}</td><td className="text-right">{$if.Out}</td></tr>);
		});
		
		return (<table className="table1 stripe-table"><thead><tr><th>Interface</th><th className="text-right nowrap" title="per second">In&nbsp;<span className="unit">ps</span></th><th className="text-right nowrap" title="per second">Out&nbsp;<span className="unit">ps</span></th><th className="text-right nowrap" title="total modulo 4G">In&nbsp;<span className="unit">%4G</span></th><th className="text-right nowrap" title="total modulo 4G">Out&nbsp;<span className="unit">%4G</span></th></tr></thead><tbody>{rows}</tbody></table>);
		
	}
});

var DFbytesCLASS = React.createClass({
	getInitialState: function() { return {DFlinks: Data.DFlinks, DFbytes: Data.DFbytes}; },

	render: function() {
		var Data = this.state;
		var rows = emptyK(Data.DFbytes, 'List') ?'': Data.DFbytes.List.map(function($disk) {
			return (<tr key={$disk.DirNameKey}><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $disk.DiskNameHTML}} /></td><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $disk.DirNameHTML}} /></td><td className="text-right">{$disk.Avail}</td><td className="text-right">{$disk.Used}&nbsp;<sup><span className={$disk.UsePercentClass}>{$disk.UsePercent}%</span></sup></td><td className="text-right">{$disk.Total}</td></tr>);
		});
		
		return (<table className="table1 stripe-table"><thead><tr><th className="header">        <a href={Data.DFlinks.DiskName.Href} className={Data.DFlinks.DiskName.Class}>Device<span  className={Data.DFlinks.DiskName.CaretClass} /></a></th><th className="header">        <a href={Data.DFlinks.DirName.Href}  className={Data.DFlinks.DirName.Class} >Mounted<span className={Data.DFlinks.DirName.CaretClass}  /></a></th><th className="header text-right"><a href={Data.DFlinks.Avail.Href}    className={Data.DFlinks.Avail.Class}   >Avail<span   className={Data.DFlinks.Avail.CaretClass}    /></a></th><th className="header text-right"><a href={Data.DFlinks.Used.Href}     className={Data.DFlinks.Used.Class}    >Used<span    className={Data.DFlinks.Used.CaretClass}     /></a></th><th className="header text-right"><a href={Data.DFlinks.Total.Href}    className={Data.DFlinks.Total.Class}   >Total<span   className={Data.DFlinks.Total.CaretClass}    /></a></th></tr></thead><tbody>{rows}</tbody></table>);
		
	}
});

var DFinodesCLASS = React.createClass({
	getInitialState: function() { return {DFlinks: Data.DFlinks, DFinodes: Data.DFinodes}; },

	render: function() {
		var Data = this.state;
		var rows = emptyK(Data.DFinodes, 'List') ?'': Data.DFinodes.List.map(function($disk) {
			return (<tr key={$disk.DirNameKey}><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $disk.DiskNameHTML}} /></td><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $disk.DirNameHTML}} /></td><td className="text-right">{$disk.Ifree}</td><td className="text-right">{$disk.Iused}&nbsp;<sup><span className={$disk.IusePercentClass}>{$disk.IusePercent}%</span></sup></td><td className="text-right">{$disk.Inodes}</td></tr>);
		});
		
		return (<table className="table1 stripe-table"><thead><tr><th className="header">Device</th><th className="header">Mounted</th><th className="header text-right">Avail</th><th className="header text-right">Used</th><th className="header text-right">Total</th></tr></thead><tbody>{rows}</tbody></table>);
		
	}
});

var MEMtableCLASS = React.createClass({
	getInitialState: function() { return Data.MEM; },

	render: function() {
		var Data = {MEM: this.state};
		var rows = Data.MEM.List.map(function($mem) {
			return (<tr key={$mem.Kind}><td>{$mem.Kind}</td><td className="text-right">{$mem.Free}</td><td className="text-right">{$mem.Used}&nbsp;<sup><span dangerouslySetInnerHTML={{__html: $mem.UsePercentHTML}} /></sup></td><td className="text-right">{$mem.Total}</td></tr>);
		});
		
		return (<table className="table1 stripe-table"><thead><tr><th></th><th className="text-right">Free</th><th className="text-right">Used</th><th className="text-right">Total</th></tr></thead><tbody>{rows}</tbody></table>);
		
	}
});

var CPUtableCLASS = React.createClass({
	getInitialState: function() { return Data.CPU; },

	render: function() {
		var Data = {CPU: this.state};
		var rows = Data.CPU.List.map(function($core) {
			return (<tr key={$core.N}><td className="text-right nowrap">{$core.N}</td><td className="text-right"><span className={$core.UserClass}>{$core.User}</span></td><td className="text-right"><span className={$core.SysClass}>{$core.Sys}</span></td><td className="text-right"><span className={$core.IdleClass}>{$core.Idle}</span></td></tr>);
		});
		
		return (<table className="table1 stripe-table"><thead><tr><th></th><th className="text-right nowrap">User<span className="unit">%</span></th><th className="text-right nowrap">Sys<span className="unit">%</span></th><th className="text-right nowrap">Idle<span className="unit">%</span></th></tr></thead><tbody>{rows}</tbody></table>);
		
	}
});

var PStableCLASS = React.createClass({
	getInitialState: function() { return Data.PStable; },

	render: function() {
		var Data = {PStable: this.state};
		var rows = Data.PStable.List.map(function($proc) {
			return (<tr key={$proc.PID}><td className="text-right">{$proc.PID}</td><td className="text-right"><span dangerouslySetInnerHTML={{__html: $proc.UserHTML}} /></td><td className="text-right">{$proc.Priority}</td><td className="text-right">{$proc.Nice}</td><td className="text-right">{$proc.Size}</td><td className="text-right">{$proc.Resident}</td><td className="text-center">{$proc.Time}</td><td className="nowrap"><span dangerouslySetInnerHTML={{__html: $proc.NameHTML}} /></td></tr>);
		});
		
		return (<table className="table2 stripe-table"><thead><tr><th className="header text-right"> <a href={Data.PStable.Links.PID.Href}      className={Data.PStable.Links.PID.Class}     >PID<span     className={Data.PStable.Links.PID.CaretClass}      /></a></th><th className="header text-right"> <a href={Data.PStable.Links.User.Href}     className={Data.PStable.Links.User.Class}    >USER<span    className={Data.PStable.Links.User.CaretClass}     /></a></th><th className="header text-right"> <a href={Data.PStable.Links.Priority.Href} className={Data.PStable.Links.Priority.Class}>PR<span      className={Data.PStable.Links.Priority.CaretClass} /></a></th><th className="header text-right"> <a href={Data.PStable.Links.Nice.Href}     className={Data.PStable.Links.Nice.Class}    >NI<span      className={Data.PStable.Links.Nice.CaretClass}     /></a></th><th className="header text-right"> <a href={Data.PStable.Links.Size.Href}     className={Data.PStable.Links.Size.Class}    >VIRT<span    className={Data.PStable.Links.Size.CaretClass}     /></a></th><th className="header text-right"> <a href={Data.PStable.Links.Resident.Href} className={Data.PStable.Links.Resident.Class}>RES<span     className={Data.PStable.Links.Resident.CaretClass} /></a></th><th className="header text-center"><a href={Data.PStable.Links.Time.Href}     className={Data.PStable.Links.Time.Class}    >TIME<span    className={Data.PStable.Links.Time.CaretClass}     /></a></th><th className="header">            <a href={Data.PStable.Links.Name.Href}     className={Data.PStable.Links.Name.Class}    >COMMAND<span className={Data.PStable.Links.Name.CaretClass}     /></a></th></tr></thead><tbody>{rows}</tbody></table>);
		
	}
});
