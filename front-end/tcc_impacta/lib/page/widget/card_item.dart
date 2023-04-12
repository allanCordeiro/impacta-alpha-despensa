import 'package:flutter/material.dart';

import '../../model/item_model.dart';
import '../../theme/theme_app.dart';

class CardItem extends StatelessWidget {
  final ItemModel itemsModel;

  const CardItem({super.key, required this.itemsModel});
  @override
  Widget build(BuildContext context) {
    return Card(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: Text(
              itemsModel.name!,
              style: style,
            ),
          ),
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: Text(
              itemsModel.creationDate!,
              style: style,
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(
              horizontal: 50,
              vertical: 10,
            ),
            child: Row(
              children: [
                Text(itemsModel.quantity!, style: style),
                // Row(
                //   children: [
                //     ElevatedButton(onPressed: () {}, child: const Text('-')),
                //   ],
                // )
              ],
            ),
          ),
        ],
      ),
    );
  }
}




// return GridView.builder(
//               itemCount: snapshot.data!.length,
//               gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
//                 crossAxisCount: 2,
//               ),
//               itemBuilder: (BuildContext context, int i) {
//                 return  CardItem(
//                   id: snapshot.data![i]['id'],
//                   name: snapshot.data![i]['name'],
//                   quantity: snapshot.data![i]['quantity'],
//                   creationDate: snapshot.data![i]['creationDate'],
//                   expirationDate: snapshot.data![i]['expirationDate'],
//                 );
//               },
//             );