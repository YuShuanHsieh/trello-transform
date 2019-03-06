import React from 'react';
import PropTypes from 'prop-types';

import { ExpandButton } from './expandButton';
import { ObjectView } from './objectView';

export class ArrayView extends React.Component {
  static propTypes = {
    array: PropTypes.arrayOf(PropTypes.oneOfType([PropTypes.object, PropTypes.string])).isRequired,
    property: PropTypes.string.isRequired,
  };

  state = {
    expand: true,
  }

  handleToggleExpand = (value) => {
    this.setState({
      expand: value,
    });
  }

  render() {
    const items = this.props.array.map((value, index) => {
      const id = `array-key-${index}`;
      return (
        <li key={id}>
          <ObjectView hideProperty property={index} object={value} />
        </li>
      );
    });
    return (
      <React.Fragment>
        <ExpandButton
          label={this.props.property}
          expand={this.state.expand}
          onToggle={this.handleToggleExpand}
        />
        {this.state.expand
          ? <ul>{items}</ul>
          : null }
      </React.Fragment>
    );
  }
}
