import 'package:flutter/material.dart';

import '../model/item_model.dart';

const String Api_Url = 'https://despensa.onrender.com/api/stock';

abstract class ApiRepository with ChangeNotifier {
  Future<ItemModel> getItems();
  Future<ItemModel> addItem();
}