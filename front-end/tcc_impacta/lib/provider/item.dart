import 'dart:math';

import 'package:flutter/material.dart';
import 'package:tcc_impacta/db/item_db.dart';

import '../model/item_model.dart';

class ItemProvider with ChangeNotifier {
  final Map<String, ItemModel> _items = {...Item_Db};

List<ItemModel> get all {
  return [..._items.values];
}

int get count {
  return _items.length;
}
ItemModel byIndex(int i) { 
  return _items.values.elementAt(i);
}
void put(ItemModel itemModel){
  if(itemModel == null) {
    return;
  }

  // Future<void> addItem(ItemModel item) async {
  //   final response = await http.post(
  //     Uri.parse(_baseUrl),
  //     body: json.encode(item.toJson()),
  //   );
  //   final responseData = json.decode(response.body);
  //   final newItem = ItemModel.fromJson(responseData);
  //   _items.add(newItem);
  // }
  final id = Random().nextDouble().toString();
  _items.putIfAbsent(id, () => ItemModel(
    id: id,
    name: itemModel.name,
    quantity: itemModel.quantity,
    creationDate: itemModel.creationDate,
    ),);
}
}