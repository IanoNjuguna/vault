#[derive(Debug, Clone, PartialEq)]
pub enum ItemType {
    Weapon,
    Armor,
    Potion,
}

#[derive(Debug, Clone)]
pub struct Item {
    pub name: String,
    pub item_type: ItemType,
    pub value: u32,
}

impl Item {
    pub fn new(name: String, item_type: ItemType, value: u32) -> Self {
        Item { name, item_type, value }
    }
}

#[derive(Debug)]
pub struct Inventory {
    pub items: Vec<Item>,
    pub max_slots: usize,
}

impl Inventory {
    pub fn new(max_slots: usize) -> Self {
        Inventory {
            items: Vec::new(),
            max_slots,
        }
    }

    pub fn add_item(&mut self, item: Item) -> Result<(), String> {
        if self.is_full() {
            return Err("Inventory is full".to_string());
        }
        self.items.push(item);
        Ok(())
    }

    pub fn remove_item(&mut self, index: usize) -> Result<Item, String> {
        if index >= self.items.len() {
            return Err(format!("Index {} is out of bounds", index));
        }
        Ok(self.items.remove(index))
    }

    pub fn get_item(&self, index: usize) -> Option<&Item> {
        self.items.get(index)
    }

    pub fn is_full(&self) -> bool {
        self.items.len() >= self.max_slots
    }

    pub fn total_value(&self) -> u32 {
        self.items.iter().map(|item| item.value).sum()
    }

    pub fn print_all_items(&self) {
        println!("=== Inventory Contents ===");
        if self.items.is_empty() {
            println!("Inventory is empty");
        } else {
            for (index, item) in self.items.iter().enumerate() {
                println!("{}: {} ({:?}) - {} gold", index + 1, item.name, item.item_type, item.value);
            }
        }
        println!("Slots used: {}/{}", self.items.len(), self.max_slots);
        println!("Total value: {} gold", self.total_value());
        println!("========================");
    }
}

fn main() {
    println!("=== RPG Inventory System Demo ===\n");

    // Create inventory with 5 slots
    let mut inventory = Inventory::new(5);

    // Create 5 different items
    let sword = Item::new("Iron Sword".to_string(), ItemType::Weapon, 100);
    let shield = Item::new("Wooden Shield".to_string(), ItemType::Armor, 50);
    let health_potion = Item::new("Health Potion".to_string(), ItemType::Potion, 25);
    let helmet = Item::new("Iron Helmet".to_string(), ItemType::Armor, 75);
    let mana_potion = Item::new("Mana Potion".to_string(), ItemType::Potion, 30);

    // Add 5 different items
    println!("Adding 5 items to inventory:");
    inventory.add_item(sword).unwrap();
    inventory.add_item(shield).unwrap();
    inventory.add_item(health_potion).unwrap();
    inventory.add_item(helmet).unwrap();
    inventory.add_item(mana_potion).unwrap();
    inventory.print_all_items();
    println!();

    // Try adding an item when inventory is full
    println!("Attempting to add an item to full inventory:");
    let extra_item = Item::new("Magic Ring".to_string(), ItemType::Weapon, 150);
    match inventory.add_item(extra_item) {
        Ok(_) => println!("Item added successfully!"),
        Err(e) => println!("Error: {}", e),
    }
    println!();

    // Remove a specific item (index 1 - Wooden Shield)
    println!("Removing item at index 1 (Wooden Shield):");
    match inventory.remove_item(1) {
        Ok(removed_item) => println!("Removed: {} ({:?})", removed_item.name, removed_item.item_type),
        Err(e) => println!("Error: {}", e),
    }
    inventory.print_all_items();
    println!();

    // Calculate and display total value
    println!("Final inventory total value: {} gold", inventory.total_value());
    println!();

    // Display all items one more time
    inventory.print_all_items();
}
