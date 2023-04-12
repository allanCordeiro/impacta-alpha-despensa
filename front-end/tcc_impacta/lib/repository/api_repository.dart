import 'package:flutter/material.dart';

import '../model/item_model.dart';

const String Api_Url = 'https://despensa.onrender.com/api/stock';

abstract class ApiRepository with ChangeNotifier {
  Future<ItemModel> getItems();
  Future<ItemModel> addItem();
}


// class Client {
//   Future<dynamic> getList() async {
//   var url = Uri.parse('https://despensa.onrender.com/api/stock');
//   var response = await http.get(url);
//   if (response.statusCode == 200) {
//     return jsonDecode(response.body);
//   } else {
//     throw Exception('Não foi possível carregar a lista');
//   }
// }

// }