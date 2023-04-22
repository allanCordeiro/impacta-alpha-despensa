import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

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
  final HttpApiRepository client = HttpApiRepository();
  @override
  Widget build(BuildContext context) {
    return FutureBuilder<dynamic>(
      future: client.getList(),
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return Scaffold(
            body: GridView.builder(
              primary: false,
              padding: const EdgeInsets.all(20),
              itemCount: snapshot.data!.length,
              gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 2,
                childAspectRatio: 4 / 3.4,
              ),
              itemBuilder: (BuildContext context, int index) {
                return Card(
                  elevation: 5,
                  child: Column(
                    children: [
                      Padding(
                        padding: const EdgeInsets.all(8.0),
                        child: Text(
                          snapshot.data[index]['name'],
                          style: const TextStyle(fontSize: 20),
                        ),
                      ),
                      Padding(
                        padding: const EdgeInsets.all(8.0),
                        child: Text(
                          snapshot.data[index]['expiration_date']
                              .split('-')
                              .reversed
                              .join('-'),
                          style: const TextStyle(fontSize: 16),
                        ),
                      ),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Row(
                            children: [
                              // IconButton(
                              //   onPressed: () async {
                              //     await client
                              //         .removeItem(snapshot.data[index]['id']);
                              //     setState(() {});
                              //   },
                              //   icon: const Icon(Icons.edit_square),
                              //   color: Colors.orange,
                              // ),
                              Padding(
                                padding:
                                    const EdgeInsets.symmetric(horizontal: 6.0),
                                child: Text(snapshot.data[index]['quantity'],
                                    style: const TextStyle(fontSize: 18)),
                              ),
                            ],
                          ),
                          IconButton(
                            onPressed: () async {
                              await client
                                  .removeItem(snapshot.data[index]['id']);
                              setState(() {});
                            },
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
            floatingActionButton: FloatingActionButton.small(
              onPressed: () {
                Navigator.of(context).pushNamed(AppRoutes.ITEMS_FORM);
              },
              child: const Icon(Icons.fact_check_outlined),
            ),
          );
        } else if (snapshot.hasError) {
          return Center(child: Text('${snapshot.error}'));
        }
        return const Center(child: CircularProgressIndicator());
      },
    );
  }
}
