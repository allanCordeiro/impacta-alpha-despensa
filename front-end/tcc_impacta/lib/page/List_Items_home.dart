import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:provider/provider.dart';

import '../model/item_model.dart';
import '../provider/item.dart';
import '../repository/implementations/http_api_repository.dart';
import '../routes/app_routes.dart';

class ListItems extends StatefulWidget {
  const ListItems({Key? key}) : super(key: key);

  @override
  _ListItemsState createState() => _ListItemsState();
}

class _ListItemsState extends State<ListItems> {
  @override
  Widget build(BuildContext context) {
    ItemProvider items = Provider.of<ItemProvider>(context);
    return Scaffold(
      appBar: AppBar(
        title: const Text('Lista de Itens'),
      ),
      body: const FutureListBuild(),
      // body: GridView.builder(
      //   itemCount: items.count,
      //   gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
      //         crossAxisCount: 2,
      //       ),
      //       itemBuilder: (ctx, i) => CardItem(itemsModel: items.byIndex(i)),
      // ),
      // floatingActionButton: FloatingActionButton(
      //   onPressed: () {
      //     Navigator.of(context).pushNamed(
      //       AppRoutes.ITEMS_FORM
      //     );
      //   },
      //   child: const Icon(Icons.add),
      // ),
    );
  }
}

class FutureListBuild extends StatefulWidget {
  const FutureListBuild({
    Key? key,
  }) : super(key: key);

  @override
  State<FutureListBuild> createState() => _FutureListBuildState();
}

class _FutureListBuildState extends State<FutureListBuild> {
  final List<ItemModel> _items = [];
  final String _baseUrl = 'https://despensa.onrender.com/api/stock';

  Future<dynamic> getList() async {
    var url = Uri.parse(_baseUrl);
    var response = await http.get(url);
    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Não foi possível carregar a lista');
    }
  }

  Future<void> addItem(ItemModel item) async {
    final response = await http.post(
      Uri.parse(_baseUrl),
      body: json.encode(item.toJson()),
    );
    final responseData = json.decode(response.body);
    final newItem = ItemModel.fromJson(responseData);
    _items.add(newItem);
  }

  final client = HttpApiRepository();
  @override
  Widget build(BuildContext context) {
    return WillPopScope(
      onWillPop: () async {
        setState(() {});
        return true;
      },
      child: FutureBuilder<dynamic>(
        future: client.getList(),
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            return Scaffold(
              body: GridView.builder(
                itemCount: snapshot.data!.length,
                gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                  crossAxisCount: 2,
                ),
                itemBuilder: (BuildContext context, int index) {
                  return Card(
                    child: Column(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: Text(snapshot.data[index]['name'],
                              style: const TextStyle(fontSize: 20)),
                        ),
                        Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: Text(snapshot.data[index]['expiration_date'],
                              style: const TextStyle(fontSize: 16)),
                        ),
                        Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            Row(
                              children: [
                                const SizedBox(
                                  width: 45,
                                  height: 35,
                                  child: ElevatedButton(
                                    onPressed: null,
                                    child: Text('-'),
                                  ),
                                ),
                                Padding(
                                  padding: const EdgeInsets.symmetric(
                                      horizontal: 8.0),
                                  child: Text(snapshot.data[index]['quantity'],
                                      style: const TextStyle(fontSize: 18)),
                                ),
                                const SizedBox(
                                  width: 45,
                                  height: 35,
                                  child: ElevatedButton(
                                    onPressed: null,
                                    child: Text('+'),
                                  ),
                                ),
                              ],
                            ),
                            IconButton(
                              onPressed: () {},
                              icon: const Icon(Icons.delete),
                              color: Colors.red,
                            )
                          ],
                        ),
                      ],
                    ),
                  );
                },
              ),
              floatingActionButton: FloatingActionButton(
                onPressed: () {
                  Navigator.of(context).pushNamed(AppRoutes.ITEMS_FORM);
                },
                child: const Icon(Icons.add),
              ),
            );
          } else if (snapshot.hasError) {
            return Center(child: Text('${snapshot.error}'));
          }
          return const Center(child: CircularProgressIndicator());
        },
      ),
    );
  }
}
