import React from 'react';
import PropTypes from 'prop-types';

import { ExpandButton } from './expandButton';
import { ArrayView } from './arrayView';

export class ObjectView extends React.Component {
  static propTypes = {
    object: PropTypes.oneOfType([PropTypes.object, PropTypes.string, PropTypes.number]).isRequired,
    property: PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired,
    hideProperty: PropTypes.bool,
  }

  static defaultProps = {
    hideProperty: false,
  }

  style = { color: '#6c6c6c' }

  state = {
    expand: true,
  }

  space = { paddingLeft: '10px' };

  handleToggleExpand = (value) => {
    this.setState({
      expand: value,
    });
  }

  render() {
    const view = [];
    let mIndex = 0;
    const { object, property, hideProperty } = this.props;
    if (typeof object === 'object') {
      view.push(
        <ExpandButton
          key={`nodeId-${mIndex += 1}`}
          label={property}
          expand={this.state.expand}
          onToggle={this.handleToggleExpand}
        />,
      );

      if (!this.state.expand) {
        return view;
      }

      Object.keys(object).forEach((value) => {
        if (Array.isArray(object[value])) {
          view.push(
            <div style={this.space} key={`nodeId-${mIndex += 1}`}>
              <ArrayView key={`nodeId-${mIndex += 1}`} property={value} array={object[value]} />
            </div>,
          );
        } else {
          view.push(
            <div style={this.space} key={`nodeId-${mIndex += 1}`}>
              <ObjectView
                property={value}
                object={object[value]}
              />
            </div>,
          );
        }
      });
    } else {
      view.push(
        <p style={this.style} key={`nodeId-${mIndex += 1}`}>
          {hideProperty ? object : `${property} : ${object}`}
        </p>,
      );
    }

    return view;
  }
}
