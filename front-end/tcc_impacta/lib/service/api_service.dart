import 'package:dio/dio.dart';

import '../model/item_model.dart';

class GithubService {
  final dio = Dio();

  void getListItems() async {
    final response = await dio.get('https://despensa.onrender.com/api/stock');
    return response.data;
  }



  void addItem(ItemModel item) async {
    final dio = Dio();
    await dio.post('https://despensa.onrender.com/api/stock',
        data: item.toJSON());
  }
  
}


  // https://despensa.onrender.com/api/stock
// Post
// {
//   "creation_date": "2023-03-22",
//   "expiration_date": "2023-04-22",
//   "name": "leite",
//   "quantity": 35
// }