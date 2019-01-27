import React from 'react';
import { Button } from '@material-ui/core';
import { ExpandMore, ExpandLess } from '@material-ui/icons';

import { ArrayView } from './arrayView';

export class ObjectView extends React.Component {
  state = {
    expand: true
  }

  expandHandler = () => {
    this.setState({
      expand: !this.state.expand
    })
  }

  render() {
    let view = [];
    let mIndex = 0;
    const { object, property } = this.props
    if (typeof object === 'object') {
      view.push(
        <Button key={`nodeId-${mIndex++}`} onClick={this.expandHandler}>
          {this.state.expand ? <ExpandLess/>:<ExpandMore />}
          {property}
        </Button>
      );

      if(!this.state.expand) {
        return view
      }

      for (const key of Object.keys(object)) {
        if (Array.isArray(object[key])) {
          view.push(<ArrayView key={`nodeId-${mIndex++}`} array={object[key]}></ArrayView>)
        } else {
          view.push(<ObjectView key={`nodeId-${mIndex++}`} property={key} object={object[key]}></ObjectView>)
        }
      }
    } else {
      view.push( <p key={`nodeId-${mIndex++}`}>{property}:{object}</p>)
    }

    return view
  }
}