import React from "react";

export class APPList extends React.Component {

  updateListItem() {
    const list = this.props.list;
    if(!list) {
      return (<li>Get data failed.</li>)
    }
    if(list.length === 0) {
      return (<li>Data has been transformed. There is no data in the list.</li>)
    }
    let index = 0;
    return (
      list.map(e => {
        return <li key={index++}>{e}</li>
      })
    )
  }

  render() {
    return (
      <ul>
        {this.updateListItem()}
      </ul>
    )
  }
}