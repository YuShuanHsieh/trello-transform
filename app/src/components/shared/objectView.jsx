import React from 'react';
import PropTypes from 'prop-types';
import { Button } from '@material-ui/core';
import { ExpandMore, ExpandLess } from '@material-ui/icons';

import { ArrayView } from './arrayView';

export class ObjectView extends React.Component {
  static propTypes = {
    object: PropTypes.oneOf([PropTypes.object, PropTypes.string, PropTypes.number]).isRequired,
    property: PropTypes.oneOf([PropTypes.string, PropTypes.number]).isRequired,
  }


  state = {
    expand: true,
  }

  expandHandler = () => {
    this.setState(pre => ({
      expand: !pre.expand,
    }));
  }

  render() {
    const view = [];
    let mIndex = 0;
    const { object, property } = this.props;
    if (typeof object === 'object') {
      view.push(
        <Button key={`nodeId-${mIndex += 1}`} onClick={this.expandHandler}>
          {this.state.expand ? <ExpandLess /> : <ExpandMore />}
          {property}
        </Button>,
      );

      if (!this.state.expand) {
        return view;
      }

      Object.keys(object).forEach((value) => {
        if (Array.isArray(object[value])) {
          view.push(<ArrayView key={`nodeId-${mIndex += 1}`} array={object[value]} />);
        } else {
          view.push(<ObjectView key={`nodeId-${mIndex += 1}`} property={value} object={object[value]} />);
        }
      });
    } else {
      view.push(
        <p key={`nodeId-${mIndex += 1}`}>
          {property}
          :
          {object}
        </p>,
      );
    }

    return view;
  }
}
