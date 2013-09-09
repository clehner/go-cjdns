package cjdns

// IpTunnel_listConnections returns a list of all current IP tunnels
func (c *Conn) IpTunnel_listConnections() (response map[string]interface{}, err error) {
	response, err = SendCmd(c, "IpTunnel_listConnections", nil)
	if err != nil {
		return
	}
	return
}

//IpTunnel_allowConnection(publicKeyOfAuthorizedNode, ip6Address=0, ip4Address=0)
//IpTunnel_connectTo(publicKeyOfNodeToConnectTo)
//IpTunnel_removeConnection(connection)
//IpTunnel_showConnection(connection)
