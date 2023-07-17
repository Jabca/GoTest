import React from 'react';
import { Menu } from 'antd';

// Define your different components
import NegativeImage from './NegativeIMage';
import LastImages from './LastImages';

import './NavigationMenu.css';

interface NavigationMenuState {
  selectedItem: string;
}

class NavigationMenu extends React.Component<{}, NavigationMenuState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      selectedItem: localStorage.getItem('selectedItem') || 'negative_image'// Set the initial selected item
    };
  }

  handleMenuItemClick = (item: any) => {
    this.setState({ selectedItem: item.key });
    localStorage.setItem('selectedItem', item.key);
  };

  render() {
    const { selectedItem } = this.state;

    let content;
    switch (selectedItem) {
      case 'negative_image':
        content = <NegativeImage />;
        break;
      case 'get_last_images':
        content = <LastImages />;
        break;
    }

    return (
      <div>
      <div className='navigation-menu'>
        <Menu
          mode="horizontal"
          selectedKeys={[selectedItem]}
          onClick={this.handleMenuItemClick}
          className="center-menu"
        >
          <Menu.Item key="negative_image">negative_image</Menu.Item>
          <Menu.Item key="get_last_images">get_last_images</Menu.Item>
        </Menu>
      </div>
        <div className="content">{content}</div>
      </div>
    );
  }
}

export default NavigationMenu;
